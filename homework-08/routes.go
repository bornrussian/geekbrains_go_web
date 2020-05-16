package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"path/filepath"
)

func (serv *Server) configureRoutes() {
	serv.mux.Route("/", func(r chi.Router) {

		// файловый сервер, который показывает статичные файлы .js и .css
		workDir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(workDir, serv.config.StaticFilesDir))
		FileServer(r, serv.config.StaticFilesURL, filesDir)

		// страницы и методы сайта
		r.Get("/", serv.wwwIndex)
		r.Get("/view", serv.wwwView)
		r.Get("/new", serv.wwwNew)
		r.Post("/delete", serv.wwwDelete)
		r.Post("/push", serv.wwwPush)
	})
}