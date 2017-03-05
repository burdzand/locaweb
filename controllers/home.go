package controllers

import (
	"github.com/andersondborges/locaweb/templates"
	"net/http"
)

// Index the main page home
func Index(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "index.html", nil)
}
