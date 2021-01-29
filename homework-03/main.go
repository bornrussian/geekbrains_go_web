//
//  Задача:
//
//  1. Создайте роут и шаблон для отображения всех постов в блоге.
//  2. Создайте роут и шаблон для просмотра конкретного поста в блоге.
//  3. Создайте роут и шаблон для редактирования и создания материала.
//

package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const STATIC_FILES_DIR = "./www/static"
const STATIC_FILES_URL = "/static/"
const TEMPLATES_FILES_DIR = "./www/templates/"
const DATABASE_DIR = "./db/"

type BlogEntry struct {
	ID string
	Autor string
	Date string
	Header string
	Content string
}

type Blog struct {
	Posts []BlogEntry
}

// Корневая страница сайта: показываем весь список постов блога
func wwwIndex(w http.ResponseWriter, r *http.Request) {
	var post BlogEntry
	var blog Blog

	// Смотрим, какие есть json-файлы в каталоге
	fileNameList, err := filepath.Glob(DATABASE_DIR + "*.json")
	if err != nil {
		log.Fatal(err)
	}

	for i:= len(fileNameList)-1; i>=0; i-- {
		fileName := fileNameList[i]
		file, _ := ioutil.ReadFile(fileName)
		_ = json.Unmarshal([]byte(file), &post)

		// Имя json-файла без пути и расширения будем использовать как ID записи блога
		re := regexp.MustCompile(`\..*`) // имя файла без расширения
		fileName = re.ReplaceAllString(fileName, "");
		re = regexp.MustCompile(`.*(\\|\/)`) // имя файла без пути
		fileName = re.ReplaceAllString(fileName, "");
		post.ID = fileName

		blog.Posts = append(blog.Posts, post)
	}

	// Читаем шаблон html страницы
	html, err := template.ParseFiles(TEMPLATES_FILES_DIR + "index.html")
	if err != nil {
		log.Fatal("Failed to parse index.html:", err)
	}

	// Вставляем в html-шаблон те данные, которые получили из json-файлов
	err = html.Execute(w, blog)
	if err != nil {
		log.Fatal("Failed to execute index.html:", err)
	}
}

// Показываем одну запись блога
func wwwView(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id := r.FormValue("id")

		// Убедимся, что id - это число
		if _, err := strconv.Atoi(id); err == nil {
			var post BlogEntry
			post.ID = id

			// Читаем json-файл с контентом блог-поста
			file, fileErr := ioutil.ReadFile(DATABASE_DIR+id+".json")
			if fileErr == nil {
				_ = json.Unmarshal([]byte(file), &post)

				// Читаем шаблон html страницы
				html, err := template.ParseFiles(TEMPLATES_FILES_DIR + "view.html")
				if err != nil {
					log.Fatal("Failed to parse view.html:", err)
				}

				// Вставляем в html-шаблон те данные, которые получили из json-файла
				err = html.Execute(w, post)
				if err != nil {
					log.Fatal("Failed to execute view.html:", err)
				}
			} else {
				// Не нашлось нужного json-файлика
				w.Write([]byte("Ahtung!"))
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

		// Убедимся, что id - это число и пробуем прочитать соответствующий json-файл
		if _, err := strconv.Atoi(id); err == nil {
			file, fileErr := ioutil.ReadFile(DATABASE_DIR+id+".json")
			if fileErr == nil {
				_ = json.Unmarshal([]byte(file), &post)
				post.ID = id
			}
		}

		// Читаем шаблон html страницы
		html, err := template.ParseFiles(TEMPLATES_FILES_DIR + "edit.html")
		if err != nil {
			log.Fatal("Failed to parse edit.html:", err)
		}

		// Вставляем в html-шаблон те данные, которые получили из json-файла
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
		id := r.PostFormValue("id")
		if id == "" {
			// Идентификатора не оказалось. Надо придумать его самостоятельно. Посмотрим на список json-файлов и придумаем новую цифру.

			// Смотрим, какие есть json-файлы в каталоге
			fileNameList, err := filepath.Glob(DATABASE_DIR + "*.json")
			if err != nil {
				log.Fatal(err)
			}
			maxFileNameID := 0
			for _, fileName := range fileNameList {
				// Имя json-файла без пути и расширения будем использовать как ID записи блога
				re := regexp.MustCompile(`\..*`) // имя файла без расширения
				fileName = re.ReplaceAllString(fileName, "");
				re = regexp.MustCompile(`.*(\\|\/)`) // имя файла без пути
				fileName = re.ReplaceAllString(fileName, "");
				if fileNameID, err := strconv.Atoi(fileName); err == nil {
					if fileNameID > maxFileNameID {
						maxFileNameID = fileNameID
					}
				}
			}
			// Пристаиваем новый ID, который на единицу больше максимального найденного среди json-файлов
			id = strconv.Itoa(maxFileNameID+1)
		}

		//Сохраняем JSON в файл:
		post := BlogEntry{
			ID: id,
			Header: r.PostFormValue("header"),
			Content: r.PostFormValue("content"),
			Autor: r.PostFormValue("autor"),
			Date: r.PostFormValue("date"),
		}
		jsonContent, _ := json.Marshal(post)
		ioutil.WriteFile(DATABASE_DIR+id+".json", jsonContent, os.ModePerm)
		http.Redirect(w, r, "http://localhost:8080", 307)
	}

}

// Удаляем запись блога
func wwwDelete(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id := r.FormValue("id")

		// Убедимся, что id - это число и пробуем прочитать соответствующий json-файл
		if _, err := strconv.Atoi(id); err == nil {
			_, fileErr := ioutil.ReadFile(DATABASE_DIR+id+".json")
			if fileErr == nil {

				// Удаляем файл
				os.Remove(DATABASE_DIR+id+".json")
			}
		}
		http.Redirect(w, r, "http://localhost:8080", 307)
	}
}

func main() {
	http.Handle(STATIC_FILES_URL, http.StripPrefix(STATIC_FILES_URL, http.FileServer(http.Dir(STATIC_FILES_DIR))))
	http.HandleFunc("/", wwwIndex)
	http.HandleFunc("/view", wwwView)
	http.HandleFunc("/edit", wwwEdit)
	http.HandleFunc("/push", wwwPush)
	http.HandleFunc("/delete", wwwDelete)
	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}