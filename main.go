package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
