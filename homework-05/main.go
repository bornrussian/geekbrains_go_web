//
//  Задача:
//
//  Перевести ваш блог на одну из ORM: gorm || beego-orm || sqlboiler (рекомендую sqlboiler)
//


package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"net/http"
	"log"
)

const STATIC_FILES_DIR = "./www/static"
const STATIC_FILES_URL = "/static/"
const TEMPLATES_FILES_DIR = "./www/templates/"
const DATABASE_LOCATION = "username:password@tcp(127.0.0.1:3306)/database"

type HasDatabase struct {
	DB *sql.DB
	CTX context.Context
}

var Site HasDatabase

func main() {
	// Пытаемся установить связь с MySQL-сервером
	db, err := sql.Open("mysql", DATABASE_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	Site.DB = db
	defer Site.DB.Close()

	ctx := context.Background()
	Site.CTX = ctx

	boil.SetDB(db)
	boil.DebugMode = true

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