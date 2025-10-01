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

	// financeRequest := types.FinanceRequest{}
	// financeRequest.Name = string([]byte(r.FormValue("name")))
	// financeRequest.Amount, _ = strconv.ParseFloat(r.FormValue("amount"), 64)
	// financeRequest.Category = r.FormValue("category")

	// if financeRequest.Amount == nil || financeRequest.Name == nil || financeRequest.Category == nil {
	// 	utils.Encode(w, &map[string]string{"error": "All fields are required!"}, http.StatusBadRequest)
	// 	return
	// }

	name := r.FormValue("name")
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	category := r.FormValue("category")

	_, err := storage.CreateFinance(name, amount, category)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	finances, _ := storage.GetFinances()
	Templates.ExecuteTemplate(w, "finances", map[string]types.Finances{"Finances": finances})
}
