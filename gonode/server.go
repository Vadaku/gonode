package main

import (
	"gonode/api"
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/api/mine", api.SendMine)
	http.HandleFunc("/api/data/", api.GetData)
	log.Fatal(http.ListenAndServe(":3222", nil))
}
