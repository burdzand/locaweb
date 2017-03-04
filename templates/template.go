package templates

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	fmt.Println("INIT TEMPLATE")
}

// RenderTemplate render a template
func RenderTemplate(w http.ResponseWriter, template string, viewModel interface{}) {
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
