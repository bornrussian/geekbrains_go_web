//
//  Задача:
//
//  1. Создайте модели для своих структур в БД.
//  2. Создайте методы для получения данных из БД по своим моделям.
//  3. Адаптируйте роуты, которые обрабатывают запросы на получение всех постов, конкретного поста в блоге и страниц редактирования.
//
//  Необходимо переделать блог на работу с БД.
//  БД можно использовать либо mysql, либо postgresql.
//

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const STATIC_FILES_DIR = "./www/static"
const STATIC_FILES_URL = "/static/"
const TEMPLATES_FILES_DIR = "./www/templates/"
const DATABASE_LOCATION = "username:password@tcp(127.0.0.1:3306)/dbname"

type BlogEntry struct {
	ID string
	Autor string
	Date string
	Header string
	Content string
}

type Blog struct {
	DB *sql.DB
}

var site Blog

func (b *Blog) GetAllJokesFromDB() ([]BlogEntry, error) {
	var res []BlogEntry

	rows, err := b.DB.Query("SELECT * FROM jokes ORDER BY date DESC")
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		joke := BlogEntry{}
		err := rows.Scan(&joke.ID, &joke.Autor, &joke.Date, &joke.Header, &joke.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		res = append(res, joke)
	}
	return res, nil
}

func (b *Blog) GetSingleJokeFromDB (id int) (BlogEntry, error) {
	var res BlogEntry

	err := b.DB.QueryRow(fmt.Sprintf("SELECT * FROM jokes WHERE id = %v", id)).
		Scan(&res.ID, &res.Autor, &res.Date, &res.Header, &res.Content)
	if err != nil {
		return res, err
	}
	return res, nil
}

func MysqlRealEscapeString(value string) string {
	replace := map[string]string{"\\":"\\\\", "'":`\'`, "\\0":"\\\\0", "\n":"\\n", "\r":"\\r", `"`:`\"`, "\x1a":"\\Z"}
	for b, a := range replace {
		value = strings.Replace(value, b, a, -1)
	}
	return value
}

func (b *Blog) PushJokeToDB (joke BlogEntry) error {
	if joke.ID == "" {
		// Если ID пустой, то вставляем новую joke в базу данных
		insert, err := b.DB.Prepare("INSERT INTO jokes (id, autor, date, header, content) VALUES (NULL, ?, ?, ?, ?);")
			if err != nil {
				log.Println(err)
			} else {
				res, err := insert.Exec(
					MysqlRealEscapeString(joke.Autor), MysqlRealEscapeString(joke.Date),
					MysqlRealEscapeString(joke.Header),	MysqlRealEscapeString(joke.Content))
				if err != nil {
					log.Println(res)
				}
				defer insert.Close()
			}
	} else {
		// Если ID не пустой, то обновляем в базе данных строчку joke с таким ID
		// Убедимся, что ID - это число
		if id, errConv := strconv.Atoi(joke.ID); errConv == nil {
			update, errUpdate := b.DB.Prepare("UPDATE jokes SET autor = ?, date = ?, header = ?, content = ? WHERE jokes.id = ?;")
			if errUpdate != nil {
				log.Println(errUpdate)
			} else {
				res, errExec := update.Exec(
					MysqlRealEscapeString(joke.Autor), MysqlRealEscapeString(joke.Date),
					MysqlRealEscapeString(joke.Header),	MysqlRealEscapeString(joke.Content), id)
				if errExec != nil {
					log.Println(res)
				}
				defer update.Close()
			}
		} else {
			return errors.New("ID is not number")
		}
	}
	return nil
}

func (b *Blog) DeleteJokeAtDB(id int) {
	delete, errDelete := b.DB.Prepare("DELETE FROM jokes WHERE id = ? LIMIT 1;")
	if errDelete != nil {
		log.Println(errDelete)
	} else {
		res, errExec := delete.Exec(id)
		if errExec != nil {
			log.Println(res)
		}
		defer delete.Close()
	}
}

// Корневая страница сайта: показываем весь список постов блога
func wwwIndex(w http.ResponseWriter, r *http.Request) {
	// Получаем записи блога из SQL-таблицы
	jokes, errJokes := site.GetAllJokesFromDB()
	if errJokes != nil {
		log.Println(errJokes)
	}

	// Готовим контент, который сунем в шаблон
	content := struct {
		Posts []BlogEntry
	}{Posts: jokes}

	// Читаем шаблон html страницы
	html, errHTML := template.ParseFiles(TEMPLATES_FILES_DIR + "index.html")
	if errHTML != nil {
		log.Fatal("Failed to parse index.html:", errHTML)
	}

	// Вставляем в html-шаблон те данные, которые получили
	errExec := html.Execute(w, content)
	if errExec != nil {
		log.Fatal("Failed to execute index.html:", errExec)
	}
}

// Показываем одну запись блога
func wwwView(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id := r.FormValue("id")

		// Убедимся, что id - это число
		if i, err := strconv.Atoi(id); err == nil {
			post, err := site.GetSingleJokeFromDB(i)
			if err == nil {
				// Читаем шаблон html страницы
				html, err := template.ParseFiles(TEMPLATES_FILES_DIR + "view.html")
				if err != nil {
					log.Fatal("Failed to parse view.html:", err)
				}

				// Вставляем в html-шаблон те данные, которые получили
				err = html.Execute(w, post)
				if err != nil {
					log.Fatal("Failed to execute view.html:", err)
				}
			}
		}
	}
}

// Редактируем запись блога
func wwwEdit(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id := r.FormValue("id")
		var post BlogEntry

		// Убедимся, что id - это число и пробуем прочитать информацию
		if i, convErr := strconv.Atoi(id); convErr == nil {
			post, _ = site.GetSingleJokeFromDB(i)
		}

		// Читаем шаблон html страницы
		html, err := template.ParseFiles(TEMPLATES_FILES_DIR + "edit.html")
		if err != nil {
			log.Fatal("Failed to parse edit.html:", err)
		}

		// Вставляем в html-шаблон те данные, которые получили
		err = html.Execute(w, post)
		if err != nil {
			log.Fatal("Failed to execute edit.html:", err)
		}
	}
}

// Вставляем запись блога
func wwwPush(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		joke := BlogEntry{
			ID: r.PostFormValue("id"),
			Header: r.PostFormValue("header"),
			Content: r.PostFormValue("content"),
			Autor: r.PostFormValue("autor"),
			Date: r.PostFormValue("date"),
		}
		site.PushJokeToDB(joke)
		http.Redirect(w, r, "http://localhost:8080", 307)
	}
}

// Удаляем запись блога
func wwwDelete(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id := r.FormValue("id")

		// Убедимся, что id - это число
		if i, err := strconv.Atoi(id); err == nil {
			site.DeleteJokeAtDB(i)
		}
		http.Redirect(w, r, "http://localhost:8080", 307)
	}
}

func main() {
	// Пытаемся установить связь с MySQL-сервером
	db, err := sql.Open("mysql", DATABASE_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	site.DB = db
	defer site.DB.Close()

	http.Handle(STATIC_FILES_URL, http.StripPrefix(STATIC_FILES_URL, http.FileServer(http.Dir(STATIC_FILES_DIR))))
	http.HandleFunc("/", wwwIndex)
	http.HandleFunc("/view", wwwView)
	http.HandleFunc("/edit", wwwEdit)
	http.HandleFunc("/push", wwwPush)
	http.HandleFunc("/delete", wwwDelete)
	log.Println("Listening on :8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}