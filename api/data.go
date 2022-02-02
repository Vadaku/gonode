package api

import (
	"gonode/database"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	datahash := r.URL.Path[13:]
	if r.Method != "GET" {
		http.Error(w, "Expected GET Request", http.StatusNotFound)
	} else {
		database.LookupInDB(datahash)
	}
}
