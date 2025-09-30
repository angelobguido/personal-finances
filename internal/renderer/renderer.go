package renderer

import (
	"html/template"
	"net/http"

	"github.com/angelobguido/personal-finances/internal/storage"
	"github.com/angelobguido/personal-finances/internal/types"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("./templates/index.html")

	finances, _ := storage.GetFinances()

	page := types.Page{Title: "Finances", Finances: finances}

	tmpl.Execute(w, page)

}
