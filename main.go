package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homePage)
	r.Post("/grid", handlerNewGrid)
	r.Get("/new", handlerNew)
	r.Post("/words/", addWord)
	r.Post("/api/words", addWordApi)
	fmt.Println("Running on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}

func returnPage(w http.ResponseWriter, templateName string) {
	basePath := "templates"
	t, err := template.ParseFiles("templates/base.html", basePath+"/"+templateName+".html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	returnPage(w, "home")
}
