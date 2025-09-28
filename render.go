package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title    string
	Finances Finances
}

func renderHome(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("./templates/index.html")

	finances, _ := getFinancesList()

	page := Page{Title: "My Home Page", Finances: finances}

	tmpl.Execute(w, page)

}
