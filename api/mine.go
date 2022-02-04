package api

import (
	"fmt"
	"gonode/functions"
	"net/http"
)

type Mine struct {
	UserID string
}

//Function for processing when the client calls the mine endpoint.
func RecieveMine(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Expected POST request", http.StatusNotFound)
	} else {
		fmt.Printf("Recieved mining request from %s\n", r.RemoteAddr)
		r.ParseForm()
		source := r.Form["source"][0]
		data := r.Form["data"][0]
		target := r.Form["target"][0]
		mode := r.Form["mode"][0]
		if mode == "fork" {
			functions.MinerFork(source, data, target)
		} else {
			functions.MinerRotate(source, data, target)
		}
	}
	// database.AddToDB()
}
