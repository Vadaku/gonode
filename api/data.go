package api

import (
	"gonode/database"
	"net/http"
	"path"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	datahash := path.Base(r.URL.Path)
	if r.Method != "GET" {
		http.Error(w, "Expected GET Request", http.StatusNotFound)
	} else {
		v := database.DBDataLookup(datahash)
		w.Write(v)
	}
}
