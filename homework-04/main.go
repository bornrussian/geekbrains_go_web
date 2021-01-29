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
	"log"
	"net/http"
)

const STATIC_FILES_DIR = "./www/static"
const STATIC_FILES_URL = "/static/"
const TEMPLATES_FILES_DIR = "./www/templates/"
const DATABASE_LOCATION = "username:password@tcp(127.0.0.1:3306)/dbname"

var site Blog

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