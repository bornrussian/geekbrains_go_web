package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"path/filepath"

	// Swagger docs for this server
	_ "geekbrains_go_web/homework-07/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (serv *Server) configureRoutes() {
	serv.mux.Route("/", func(r chi.Router) {

		// файловый сервер, который показывает статичные файлы .js и .css
		workDir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(workDir, STATIC_FILES_DIR))
		FileServer(r, STATIC_FILES_URL, filesDir)

		// страницы и методы сайта
		r.Get("/", serv.wwwIndex)
		r.Get("/view", serv.wwwView)
		r.Get("/new", serv.wwwNew)
		r.Get("/edit", serv.wwwEdit)
		r.Get("/delete", serv.wwwDelete)
		r.Post("/push", serv.wwwPush)

		// документация кода проекта
		r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(serv.swagURL)))
		r.Get("/"+serv.swagPath, serv.HandleSwagger)
	})
}