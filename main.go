package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/tempe", TempeHandler)
	http.HandleFunc("/length", LengthHandler)
	http.HandleFunc("/weigth", WeigthHandler)
	http.ListenAndServe(":3000", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}
func LengthHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/length.html"))
	tmpl.Execute(w, nil)
}

func WeigthHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/weigth.html"))
	tmpl.Execute(w, nil)
}

func TempeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/tempe.html"))
	tmpl.Execute(w, nil)
}
