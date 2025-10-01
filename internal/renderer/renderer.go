package renderer

import (
	"html/template"
	"net/http"

	"github.com/angelobguido/personal-finances/internal/storage"
	"github.com/angelobguido/personal-finances/internal/types"
)

var Templates *template.Template

func RenderHome(w http.ResponseWriter, r *http.Request) {

	finances, _ := storage.GetFinances()

	page := types.Page{Title: "Finances", Finances: finances}

	Templates.ExecuteTemplate(w, "index", page)

}
