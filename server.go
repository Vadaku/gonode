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
	http.HandleFunc("/api/v2/index/", api.GetIndex)
	http.HandleFunc("/auth/callback", api.SpotifyCallback)
	http.HandleFunc("/auth/login", api.SpotifyLogin)
	log.Fatal(http.ListenAndServe(":3222", nil))

}
