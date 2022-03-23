package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Setup routes and handlers then serve on port 8080.
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v2/mine", mine).Methods("POST")
	r.HandleFunc("/api/v2/hashwall", hashwall).Methods("PUT")
	r.HandleFunc("/api/v2/data", getData).Methods("GET")
	r.HandleFunc("/api/v2/index", getIndex).Methods("GET")

	fmt.Println("Running server on port 3222.")
	http.ListenAndServe(":3222", r)
}
