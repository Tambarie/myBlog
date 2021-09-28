package main

import (
	"github.com/Tambarie/myBlog/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"html/template"
	"log"
	"net/http"
	"time"
)

func (app *application) homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//not found error
		app.notFound(w)
		return
	}
	templateFiles := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
		"./ui/html/form.page.html",
	}
	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		//server error
		app.serverError(w, err)
		return
	}

	var blog = models.Blog{}
	users, err := blog.GetAllContents()
	if err != nil{
		app.errorLog.Println("Can't read from the database")
	}

	err = ts.Execute(w, users)
	if err != nil {
		//server error
		app.serverError(w, err)
	}

}

func (app *application) getFormPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/formPage" {
		app.notFound(w)
		return
	}

	templateFiles := []string{
		"./ui/html/form.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/home.page.html",
		"./ui/html/footer.partial.html",
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

func (app *application) postForm(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/postForm"{
		app.notFound(w)
		return
	}
	author := r.PostFormValue("author")
	title := r.PostFormValue("title")
	body := r.PostFormValue("article")

	var blog = &models.Blog{
		Author: author,
		Title: title,
		Body: body,
		Time: time.Now().Format("2006-January-02"),
		ID: uuid.New().String(),
	}

	blog.CreateBlog()

	http.Redirect(w,r,"/",http.StatusFound)
}

func (app *application) Delete(w http.ResponseWriter, r *http.Request)  {
	blog := models.Blog{}
	id := chi.URLParam(r,"id")

	err := blog.Delete(id)
	app.infoLog.Printf(id)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w,r,"/",http.StatusFound)
	//71697cc7-c919-4d0d-b3ad-97f6fc6b51b2


}

func (app *application) EditPage(w http.ResponseWriter, r *http.Request)  {
	templateFiles := []string{
		"./ui/html/edit.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/home.page.html",
		"./ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(templateFiles...)
	if err != nil{
		panic(err)
	}

	blog := &models.Blog{}
	id := chi.URLParam(r,"id")
	data := blog.GetContent(id)

	if err := ts.Execute(w,data); err != nil{
		 app.errorLog.Printf("could not execute file")
	}

}

func (app *application) upDatePost(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")

	author := r.PostFormValue("author")
	title := r.PostFormValue("title")
	body := r.PostFormValue("article")

	var blog = &models.Blog{
		ID: id,
		Author: author,
		Title: title,
		Body: body,
	}

	log.Println(blog)
	err := blog.UpdateForm()

	if err != nil{
		app.errorLog.Fatal("update operation failed")
	}
	http.Redirect(w,r,"/",http.StatusFound)
}


