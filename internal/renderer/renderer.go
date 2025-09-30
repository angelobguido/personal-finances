package renderer

import (
	"html/template"
	"net/http"

	"github.com/angelobguido/personal-finances/internal/api"
	"github.com/angelobguido/personal-finances/internal/types"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("./templates/index.html")

	finances, _ := api.GetFinancesList()

	page := types.Page{Title: "My Home Page", Finances: finances}

	tmpl.Execute(w, page)

}
