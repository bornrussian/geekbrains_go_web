package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"log"
	"net/http"
)

// Корневая страница сайта: показываем весь список постов блога
func wwwIndex(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	// Получаем записи блога из SQL-таблицы
	jokes, errJokes := GetPosts(context.TODO(), db)
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

// Показываем одну запись блога
func wwwView(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id, _ := primitive.ObjectIDFromHex(r.FormValue("id"))
		post, errPost := GetPost(context.TODO(), db, id)
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

// Вставляем новую запись блога
func wwwNew(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
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

// Редактируем запись блога
func wwwEdit(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id, _ := primitive.ObjectIDFromHex(r.FormValue("id"))
		post, errPost := GetPost(context.TODO(), db, id)
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

// Вставляем запись блога
func wwwPush(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
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
			joke.Insert(context.TODO(), db)
		} else {
			joke.Update(context.TODO(), db)
		}
		http.Redirect(w, r, "http://localhost:8080", http.StatusTemporaryRedirect)
	}
}

// Удаляем запись блога
func wwwDelete(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	if isParsed := r.ParseForm(); isParsed != nil {
		w.Write([]byte("Something is wrong while html form parsing"))
	} else {
		id, _ := primitive.ObjectIDFromHex(r.FormValue("id"))
		postDelete := Post{
			Mongo: Mongo{
				ID: id,
			},
		}
		postDelete.Delete(context.TODO(), db)

		http.Redirect(w, r, "http://localhost:8080", http.StatusTemporaryRedirect)
	}
}
