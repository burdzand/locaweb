package controllers

import (
	"fmt"
	"net/http"
)

// ExportMostMentions export all tweets mentions
func ExportMostMentions(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "ExportMostMentions")
}
