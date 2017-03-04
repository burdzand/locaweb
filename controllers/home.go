package controllers

import (
	"github.com/andersondborges/locaweb/templates"
	"github.com/andersondborges/locaweb/twitter"
	"net/http"
)

// Index the main page home
func Index(w http.ResponseWriter, r *http.Request) {
	statuses, _ := twitter.GetTweets()
	templates.RenderTemplate(w, "index.html", statuses)
}
