package renderer

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/angelobguido/personal-finances/internal/storage"
	"github.com/angelobguido/personal-finances/internal/types"
)

var Templates *template.Template

func RenderHome(w http.ResponseWriter, r *http.Request) {

	finances, _ := storage.GetFinances()

	page := types.Page{Title: "Finances", Finances: finances}

	Templates.ExecuteTemplate(w, "index", page)

}

func CreateFinance(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	category := r.FormValue("category")

	finance, err := storage.CreateFinance(name, amount, category)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Templates.ExecuteTemplate(w, "finance-list-item", *finance)
}

func UpdateFinance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("Id")
	name := r.FormValue("name")
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	category := r.FormValue("category")

	finance, err := storage.UpdateFinanceById(id, &name, &amount, &category)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Templates.ExecuteTemplate(w, "finance-list-item", *finance)
}

func RenderEditFinance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("Id")

	finance, err := storage.GetFinanceById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Templates.ExecuteTemplate(w, "edit-finance-list-item", *finance)
}

func RenderFinance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("Id")

	finance, err := storage.GetFinanceById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Templates.ExecuteTemplate(w, "finance-list-item", *finance)
}
