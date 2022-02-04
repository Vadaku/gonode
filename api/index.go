package api

import (
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	// root := path.Base(r.URL.Path)
	// fmt.Println(root)
	if r.Method != "GET" {
		http.Error(w, "Expected GET request", http.StatusNotFound)
	} else {

		// data, err := database.DBIndexLookup(root)
	}
}
