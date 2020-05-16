package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"net/http"
	"time"
)

// @Description wwwIndex - показываем корневую страницу сайта
// @Tags handlers
// @Router / [get]
func (serv *Server) wwwIndex(w http.ResponseWriter, r *http.Request) {
	// Получаем записи блога из SQL-таблицы
	jokes, errJokes := GetPosts(context.TODO(), serv.db)
	if errJokes != nil {
		serv.lg.Error(errJokes)
	}

	// Готовим контент, который сунем в шаблон
	content := struct {
		Posts []Post
	}{Posts: jokes}

	// Читаем шаблон html страницы
	html, errHTML := template.ParseFiles(serv.config.TemplateFilesDir + "index.html")
	if errHTML != nil {
		serv.lg.Error("failed to parse index.html:", errHTML)
		return
	}

	// Вставляем в html-шаблон те данные, которые получили
	errExec := html.Execute(w, content)
	if errExec != nil {
		serv.lg.Error("failed to execute index.html:", errHTML)
		return
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
			html, errHTML := template.ParseFiles(serv.config.TemplateFilesDir + "view.html")
			if errHTML != nil {
				serv.lg.Error("failed to parse view.html:", errHTML)
				return
			}
			// Вставляем в html-шаблон те данные, которые получили
			errExec := html.Execute(w, post)
			if errExec != nil {
				serv.lg.Error("failed to execute view.html:", errHTML)
				return
			}
		} else {
			serv.lg.Error("could not find post id", r.FormValue("id"))
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
		html, errHTML := template.ParseFiles(serv.config.TemplateFilesDir + "new.html")
		if errHTML != nil {
			serv.lg.Error("failed to parse new.html:", errHTML)
			return
		}

		// Вставляем в html-шаблон те данные, которые получили
		errExec := html.Execute(w, nil)
		if errExec != nil {
			serv.lg.Error("failed to execute new.html:", errHTML)
			return
		}
	}
}

func (serv *Server) wwwPush(w http.ResponseWriter, r *http.Request) {
	ip := GetIP(r)
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		if r.PostFormValue("pin") == serv.config.PinCodeForUpload {

			joke := Post{
				Header:  r.PostFormValue("header"),
				Content: r.PostFormValue("content"),
				Autor:   ip,
				Date:    time.Now().Format("01-02-2006 15:04:05"),
			}
			joke.Insert(context.TODO(), serv.db)
			serv.lg.Info("wwwPush: new joke from ", ip)
			http.Redirect(w, r, serv.config.URL, 301)
		} else {
			w.Write([]byte("wrong pin"))
			serv.lg.Info("wwwPush: wrong pin from ", ip)
		}
	}
}

func (serv *Server) wwwDelete(w http.ResponseWriter, r *http.Request) {
	ip := GetIP(r)
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		if r.PostFormValue("pin") == serv.config.PinCodeForDelete {
			id, _ := primitive.ObjectIDFromHex(r.PostFormValue("id"))
			postDelete := Post{
				Mongo: Mongo{
					ID: id,
				},
			}
			postDelete.Delete(context.TODO(), serv.db)
			serv.lg.Info("wwwDelete: from ", ip)
			http.Redirect(w, r, serv.config.URL, 301)
		} else {
			w.Write([]byte("wrong pin"))
			serv.lg.Info("wwwDelete: wrong pin from ", ip)
		}
	}
}
