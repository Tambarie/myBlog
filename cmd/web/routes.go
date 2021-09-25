package main

import "net/http"

func (app *application) routes() *http.ServeMux  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homepage)
	mux.HandleFunc("/createForm", app.create)

	staticFileServer := http.FileServer(http.Dir("../ui/static"))
	mux.Handle("/static/",http.StripPrefix("/static",staticFileServer))

	return mux
}