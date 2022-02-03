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
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		response := database.DBDataLookup(datahash)
		w.Write(response)
	}
}
