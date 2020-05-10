//
//  Задача:
//
//  Переведите ваш блог на MongoDB.
//

package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
	"context"
)

const STATIC_FILES_DIR = "./www/static"
const STATIC_FILES_URL = "/static/"
const TEMPLATES_FILES_DIR = "./www/templates/"

const DATABASE_LOCATION = "mongodb://localhost:27017"
const DATABASE_DBNAME = "geekbrains"
const DATABASE_COLLECTION = "jokes"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_LOCATION))
	if err != nil {
		log.Fatal(err)
	}
	mongoDB := client.Database(DATABASE_DBNAME)

	http.Handle(STATIC_FILES_URL, http.StripPrefix(STATIC_FILES_URL, http.FileServer(http.Dir(STATIC_FILES_DIR))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { wwwIndex(mongoDB, w, r) })
	http.HandleFunc("/view", func(w http.ResponseWriter, r *http.Request) { wwwView(mongoDB, w, r) })
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) { wwwNew(mongoDB, w, r) })
	http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) { wwwEdit(mongoDB, w, r) })
	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) { wwwPush(mongoDB, w, r) })
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) { wwwDelete(mongoDB, w, r) })
	log.Println("Listening on :8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
