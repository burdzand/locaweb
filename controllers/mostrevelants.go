package controllers

import (
	"github.com/andersondborges/locaweb/templates"
	"github.com/andersondborges/locaweb/twitter"
	"net/http"
)

// ExportMostRelevants export all relevants tweets
func ExportMostRelevants(w http.ResponseWriter, r *http.Request) {
	tweets, _ := twitter.GetTweets()
	templates.RenderTemplate(w, "relevants.html", tweets)
}
