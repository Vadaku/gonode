package main

import (
	"gonode/api"
	"gonode/database"
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/api/v2/mine", api.RecieveMine)
	http.HandleFunc("/api/v2/data/", api.GetData)
	http.HandleFunc("/api/v2/hashwall/", api.HashWall)
	database.ConnectToDB()
	log.Fatal(http.ListenAndServe(":3222", nil))

}
