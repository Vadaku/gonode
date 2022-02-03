package main

import (
	"gonode/api"
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/api/v2/mine", api.RecieveMine)
	http.HandleFunc("/api/v2/data/", api.GetData)
	http.HandleFunc("/api/v2/hashwall/", api.HashWall)
	log.Fatal(http.ListenAndServe(":3222", nil))

}
