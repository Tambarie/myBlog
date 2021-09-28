package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Get("/", app.homepage)
	//mux.Get("/",app.getAllContents)
	mux.Get("/formPage", app.getFormPage)
	mux.Get("/{id}/delete", app.Delete)
	mux.Get("/{id}/edit",app.EditPage)
	mux.Post("/postForm", app.postForm)
	mux.Post("/{id}/updatePost",app.upDatePost)


	staticFileServer := http.FileServer(http.Dir("../ui/static"))
	mux.Handle("/static/",http.StripPrefix("/static",staticFileServer))

	return mux
}