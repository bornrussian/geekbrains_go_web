package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"log"
	"net/http"
)

// @Description wwwIndex - показываем корневую страницу сайта
// @Tags handlers
// @Router / [get]
func (serv *Server) wwwIndex(w http.ResponseWriter, r *http.Request) {
	// Получаем записи блога из SQL-таблицы
	jokes, errJokes := GetPosts(context.TODO(), serv.db)
	if errJokes != nil {
		log.Println(errJokes)
	}

	// Готовим контент, который сунем в шаблон
	content := struct {
		Posts []Post
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

// @Description wwwView - отображает одну конкретную запись блога
// @Tags handlers
// @Router /view [get]
func (serv *Server) wwwView(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id, _ := primitive.ObjectIDFromHex(r.FormValue("id"))
		post, errPost := GetPost(context.TODO(), serv.db, id)
		if errPost == nil {
			// Читаем шаблон html страницы
			html, errHTML := template.ParseFiles(TEMPLATES_FILES_DIR + "view.html")
			if errHTML != nil {
				log.Fatal("Failed to parse view.html:", errHTML)
			}
			// Вставляем в html-шаблон те данные, которые получили
			errExec := html.Execute(w, post)
			if errExec != nil {
				log.Fatal("Failed to execute view.html:", errExec)
			}
		} else {
			log.Println("Could not find post id", r.FormValue("id"))
		}
	}
}

// @Description wwwNew - показываем форму, в которой можно внести данные для нового поста блога
// @Tags handlers
// @Router /new [get]
func (serv *Server) wwwNew(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {

		// Читаем шаблон html страницы
		html, errHTML := template.ParseFiles(TEMPLATES_FILES_DIR + "new.html")
		if errHTML != nil {
			log.Fatal("Failed to parse new.html:", errHTML)
		}

		// Вставляем в html-шаблон те данные, которые получили
		errExec := html.Execute(w, nil)
		if errExec != nil {
			log.Fatal("Failed to execute new.html:", errExec)
		}
	}
}

// @Description wwwEdit - форма редактирования записи блога
// @Tags handlers
// @Router /edit [get]
func (serv *Server) wwwEdit(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id, _ := primitive.ObjectIDFromHex(r.FormValue("id"))
		post, errPost := GetPost(context.TODO(), serv.db, id)
		if errPost == nil {
			// Читаем шаблон html страницы
			html, errHTML := template.ParseFiles(TEMPLATES_FILES_DIR + "edit.html")
			if errHTML != nil {
				log.Fatal("Failed to parse edit.html:", errHTML)
			}

			// Вставляем в html-шаблон те данные, которые получили
			errExec := html.Execute(w, post)
			if errExec != nil {
				log.Fatal("Failed to execute edit.html:", errExec)
			}
		} else {
			log.Println("Could not find post id", r.FormValue("id"))
		}
	}
}

// @Description wwwPush - вставляем запись блога в базу данных
// @Tags handlers
// @Router /push [post]
func (serv *Server) wwwPush(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id, _ := primitive.ObjectIDFromHex(r.PostFormValue("id"))
		joke := Post{
			Mongo: Mongo{ ID: id },
			Header:  r.PostFormValue("header"),
			Content: r.PostFormValue("content"),
			Autor:   r.PostFormValue("autor"),
			Date:    r.PostFormValue("date"),
		}

		if r.PostFormValue("id") == "" {
			joke.Insert(context.TODO(), serv.db)
		} else {
			joke.Update(context.TODO(), serv.db)
		}
		w.Write([]byte("ok"))
	}
}

// @Description wwwDelete - удаляем запись блога из базы данных
// @Tags handlers
// @Router /delete [get]
func (serv *Server) wwwDelete(w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id, _ := primitive.ObjectIDFromHex(r.FormValue("id"))
		postDelete := Post{
			Mongo: Mongo{
				ID: id,
			},
		}
		postDelete.Delete(context.TODO(), serv.db)

		w.Write([]byte("ok"))
	}
}

// HandleSwagger - Returns swagger.json docs
// @Description Returns swagger.json docs
// @Tags system
// @Success 200 {string} string
// @Router /docs/swagger.json [get]
func (serv *Server) HandleSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, serv.swagPath)
}