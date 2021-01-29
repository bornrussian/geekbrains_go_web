package main

import (
	"fmt"
	"geekbrains_go_web/homework-05/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Корневая страница сайта: показываем весь список постов блога
func wwwIndex(w http.ResponseWriter, r *http.Request) {
	// Получаем записи блога из базы данных
	jokes, jokesError := models.Jokes().All(Site.CTX, Site.DB)
	if jokesError != nil {
		log.Println(jokesError)
	}

	// Читаем шаблон html страницы
	html, htmlError := template.ParseFiles(TEMPLATES_FILES_DIR + "index.html")
	if htmlError != nil {
		log.Fatal("Failed to parse index.html:", htmlError)
	}

	// Вставляем в html-шаблон те данные, которые получили из базы данных
	execError := html.Execute(w, struct {Posts models.JokeSlice}{Posts: jokes})
	if execError != nil {
		log.Fatal("Failed to execute index.html:", execError)
	}
}

// Показываем одну запись блога
func wwwView(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
			joke, jokeError := models.FindJoke(Site.CTX, Site.DB, int64(id))
			if jokeError == nil {
				// Читаем шаблон html страницы
				html, htmlError := template.ParseFiles(TEMPLATES_FILES_DIR + "view.html")
				if htmlError != nil {
					log.Fatal("Failed to parse view.html:", htmlError)
				}

				// Вставляем в html-шаблон те данные, которые получили
				execError := html.Execute(w, joke)
				if execError != nil {
					log.Fatal("Failed to execute view.html:", execError)
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

		// Читаем шаблон html страницы
		html, err := template.ParseFiles(TEMPLATES_FILES_DIR + "edit.html")
		if err != nil {
			log.Fatal("Failed to parse edit.html:", err)
		}

		// Убедимся, что id - это число и пробуем прочитать информацию
		if i, err := strconv.Atoi(id); err == nil {
			joke, jokeErr:= models.FindJoke(Site.CTX, Site.DB, int64(i))
			if jokeErr == nil {
				// Вставляем в html-шаблон те данные, которые получили из базы данных
				execError := html.Execute(w, joke)
				if execError != nil {
					log.Fatal("Failed to execute edit.html:", execError)
				}
				return
			}
		}
		// Вставляем пустые данные
		execError := html.Execute(w, nil)
		if execError != nil {
			log.Fatal("Failed to execute edit.html:", execError)
		}
	}
}

// Вставляем запись блога
func wwwPush(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id, errConv := strconv.Atoi(r.PostFormValue("id"))
		joke := models.Joke{
			ID: int64(id),
			Header: r.PostFormValue("header"),
			Content: r.PostFormValue("content"),
			Autor: r.PostFormValue("autor"),
			Date: r.PostFormValue("date"),
		}
		if errConv == nil {
			// если id - это цифра, то обновляем joke с таким идентификатором
			_, err := joke.Update(Site.CTX,Site.DB,boil.Infer())
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// если id - НЕ цифра, то вставляем новый joke
			err := joke.Insert(Site.CTX,Site.DB,boil.Infer())
			if err != nil {
				fmt.Println(err)
			}
		}
		http.Redirect(w, r, "http://localhost:8080", 307)
	}
}

// Удаляем запись блога
func wwwDelete(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		// Убедимся, что id - это число
		if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
			joke, jokeError := models.FindJoke(Site.CTX, Site.DB, int64(id))
			if jokeError == nil {
				joke.Delete(Site.CTX, Site.DB)
			}
		}
		http.Redirect(w, r, "http://localhost:8080", 307)
	}
}