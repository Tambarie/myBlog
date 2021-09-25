package main

import (
	"html/template"
	"net/http"
)

func (app *application) homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//not found error
		app.notFound(w)
		return
	}
	templateFiles := []string{
		"../../ui/html/home.page.html",
		"../../ui/html/base.layout.html",
		"../../ui/html/footer.partial.html",
		"../../ui/html/create.page.html",
	}
	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		//server error
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		//server error
		app.serverError(w, err)
	}
}

func (app *application) create(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/createForm" {
		app.notFound(w)
		return
	}

	templateFiles := []string{
		"../../ui/html/create.page.html",
		"../../ui/html/base.layout.html",
		"../../ui/html/home.page.html",
		"../../ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		app.serverError(w, err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}
