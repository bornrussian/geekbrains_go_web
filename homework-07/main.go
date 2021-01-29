//
//  Задача:
//
//  1. Напишите тесты для функций, с помощью которых вы работаете с базой данных.
//	2. Напишите http-тесты для методов сайта.
//	3. Добавьте документацию в проект.
//

package main

import (
	"context"
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

const STATIC_FILES_DIR = "www/static"
const STATIC_FILES_URL = "/static"
const TEMPLATES_FILES_DIR = "./www/templates/"

const DATABASE_LOCATION = "mongodb://localhost:27017"
const DATABASE_DBNAME = "geekbrains"
const DATABASE_COLLECTION = "jokes"

func main() {
	// ARGs
	flagAddr := flag.String("addr", "localhost:8080", "server address")
	flagSwagURLAddr := flag.String("swag-url", "http://localhost:8080/docs/swagger.json", "swagger.json http address")
	flagSwagJSONPath := flag.String("swag-path", "docs/swagger.json", "swagger.json path")
	flag.Parse()

	// HTTP START
	lg := logrus.New()
	serv := NewServer(context.TODO(), lg)
	serv.SetSwagger(*flagSwagJSONPath, *flagSwagURLAddr)
	serv.Start(*flagAddr)
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt, os.Kill)
	<-stopChan
	serv.Stop()
}
