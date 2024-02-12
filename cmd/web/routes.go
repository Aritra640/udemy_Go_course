package main

import (
	"net/http"

	handlers "github.com/Aritra640/Go_course/HTMLdemo/pkg"
	"github.com/Aritra640/Go_course/HTMLdemo/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)
	mux.Use(WriteToConsole)
	mux.Use(NoCSRF)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
} /* */
