package api

import (
	"gonode/database"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	database.DBListData()
	if r.Method != "GET" {
		http.Error(w, "Expected GET Request", http.StatusNotFound)
	} else {
	}
}
