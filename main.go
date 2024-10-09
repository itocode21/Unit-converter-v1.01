package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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

func convertLenght(value float64, fromUnit, toUnit string) (float64, error) {
	switch fromUnit + "to" + toUnit {
	case "millimeter to centimeter":
		return (value / 10.0), nil
	case "millimeter to meter":
		return (value / 1000.0), nil
	case "millimeter to inch":
		return (value / 25.4), nil
	case "millimeter to foot":
		return (value / 304.8), nil
	case "millmeter to yard":
		return (value / 914.4), nil
	case "millimeter to mile":
		return (value / 1609344.0), nil
	default:
		return value, fmt.Errorf("not supporte from %s to %s", fromUnit, toUnit)
	}

}

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		value, _ := strconv.ParseFloat(r.FormValue("value"), 64)
		formUnit := r.FormValue("FromUnit")
		toUnit := r.FormValue("toUnit")
		result, _ := convertLenght(value, formUnit, toUnit)
		tmpl := template.Must(template.ParseFiles("templates/length.html"))
		tmpl.Execute(w, map[string]interface{}{
			"Result": fmt.Sprintf("%.2f", result),
		})
		return
	}

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
