package controllers

import (
	"encoding/json"
	"github.com/andersondborges/locaweb/templates"
	"github.com/andersondborges/locaweb/twitter"
	"net/http"
	"sort"
)

// ControllerMostRelevants show and export all relevants tweets
func ControllerMostRelevants(w http.ResponseWriter, r *http.Request) {
	tweets, err := twitter.GetTweets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sort.Sort(twitter.SortByFavorite(tweets))
	sort.Sort(twitter.SortByRetweets(tweets))
	sort.Sort(twitter.SortByFollows(tweets))

	export := r.URL.Query().Get("export")
	if export == "json" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tweets)
	} else {
		templates.RenderTemplate(w, "relevants.html", tweets)
	}
}
