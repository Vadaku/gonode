package api

import (
	"fmt"
	"net/http"
)

func SendMine(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path[10:])
	// if r.Method != "POST" {
	// 	http.Error(w, "Expected POST request", http.StatusNotFound)
	// }
}
