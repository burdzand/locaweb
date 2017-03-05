package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var tmpl *template.Template

func init() {

	//	filepath.Join(os.Getenv("GOPATH"), "src/github.com/andersondborges/locaweb/templates/*")

	tmpl = template.Must(template.ParseGlob(filepath.Join(os.Getenv("GOPATH"), "src/github.com/andersondborges/locaweb/templates/*")))
	fmt.Println("INIT TEMPLATE")
}

// RenderTemplate render a template
func RenderTemplate(w http.ResponseWriter, template string, viewModel interface{}) {
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
