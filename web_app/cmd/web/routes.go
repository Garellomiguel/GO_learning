package main

import (
	"net/http"
	"web_app/pkg/config"
	"web_app/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// Set were to go look for local file to use on tempaltes
	// fileServer := http.FileServer(http.Dir("./img/"))
	// mux.Handle("/img/*", http.StripPrefix("/img", fileServer))

	return mux
}
