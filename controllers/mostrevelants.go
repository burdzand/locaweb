package controllers

import (
	"encoding/json"
	"github.com/andersondborges/locaweb/templates"
	"github.com/andersondborges/locaweb/twitter"
	"net/http"
	"sort"
)

// ExportMostRelevants export all relevants tweets
func ExportMostRelevants(w http.ResponseWriter, r *http.Request) {
	tweets, _ := twitter.GetTweets()
	sort.Sort(twitter.SortByFavorite(tweets))
	sort.Sort(twitter.SortByRetweets(tweets))
	sort.Sort(twitter.SortByFollows(tweets))

	export := r.URL.Query().Get("export")
	if export == "json" {
		json.NewEncoder(w).Encode(tweets)
	} else {
		templates.RenderTemplate(w, "relevants.html", tweets)
	}
}
