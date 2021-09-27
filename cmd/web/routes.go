package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()
	//mux := http.NewServeMux()
	mux.HandleFunc("/", app.homepage)
	mux.HandleFunc("/createForm", app.create)

	staticFileServer := http.FileServer(http.Dir("../ui/static"))
	mux.Handle("/static/",http.StripPrefix("/static",staticFileServer))

	return mux
}