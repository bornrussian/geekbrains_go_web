package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

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