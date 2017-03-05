package controllers

import (
	"encoding/json"
	"github.com/andersondborges/locaweb/templates"
	"github.com/andersondborges/locaweb/twitter"
	"net/http"
	"sort"
)

// ControllerMostMentions show and export all tweets mentions
func ControllerMostMentions(w http.ResponseWriter, r *http.Request) {

	tusers := twitter.GetTweetsOrderUser()
	if tusers == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sort.Sort(twitter.SortUserTweetsByFollows(tusers))
	export := r.URL.Query().Get("export")
	if export == "json" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tusers)
	} else {
		templates.RenderTemplate(w, "usermentions.html", tusers)
	}
}
