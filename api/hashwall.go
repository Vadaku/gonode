package api

import (
	"net/http"
)

func HashWall(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Expected GET request", http.StatusNotFound)
	} else {

	}
}
