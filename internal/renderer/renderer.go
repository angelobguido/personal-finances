package renderer

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/angelobguido/personal-finances/internal/storage"
	"github.com/angelobguido/personal-finances/internal/types"
	"github.com/angelobguido/personal-finances/internal/utils"
)

var Templates *template.Template

func RenderHome(w http.ResponseWriter, r *http.Request) {

	finances, _ := storage.GetFinances()

	page := types.Page{Title: "Finances", Finances: finances}

	Templates.ExecuteTemplate(w, "finances.html", page)

}

func CreateFinance(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		utils.Encode(w, &map[string]string{"error": "The amount value is wrong!"}, http.StatusBadRequest)
		return
	}
	category := r.FormValue("category")
	createdAt, err := time.Parse("2006-01-02T15:04", r.FormValue("created_at"))
	if err != nil {
		utils.Encode(w, &map[string]string{"error": "The created_at value is wrong!"}, http.StatusBadRequest)
		return
	}

	finance, err := storage.CreateFinance(name, amount, category, createdAt)

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
	createdAt, _ := time.Parse("2006-01-02T15:04", r.FormValue("created_at"))

	finance, err := storage.UpdateFinanceById(id, utils.PtrIfNotZero(name), utils.PtrIfNotZero(amount), utils.PtrIfNotZero(category), utils.PtrIfNotZero(createdAt))

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
