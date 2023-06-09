package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type Usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		//	templates.ExecuteTemplate(w, "html.html", nil)
		u := Usuario{
			"Maria",
			"joao@gmail.com",
		}

		templates.ExecuteTemplate(w, "html.html", u)

	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
