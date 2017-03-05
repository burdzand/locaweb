package controllers

import (
	"encoding/json"
	"github.com/andersondborges/locaweb/templates"
	"github.com/andersondborges/locaweb/twitter"
	"net/http"
	"sort"
)

// ExportMostMentions export all tweets mentions
func ExportMostMentions(w http.ResponseWriter, r *http.Request) {

	tusers := twitter.GetTweetsOrderUser()
	sort.Sort(twitter.SortUserTweetsByFollows(tusers))
	export := r.URL.Query().Get("export")
	if export == "json" {
		json.NewEncoder(w).Encode(tusers)
	} else {
		templates.RenderTemplate(w, "usermentions.html", tusers)
	}
}
