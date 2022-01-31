package api

import (
	"fmt"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path[10:])
	if r.Method != "GET" {
		http.Error(w, "Expected GET Request", http.StatusNotFound)
	}
}
