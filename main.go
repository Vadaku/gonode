package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var test *Trie

//Setup routes and handlers then serve on port 8080.
func main() {
	//Init Imgui.
	// util.InitImgui()
	// SetupRoutes()
	//Init and test Trie.
	test = initializeTrie()
	test.insertToTrie("21e8", "21e893411ac5c7f3896fe57fb7d8a8f150ee18a7256fe73990a17c47a498c8b5")
	test.insertToTrie("21e8", "21eb2f005c551eca25903ab09dbd08f512d9cbb6af226152690583cbcac51135")
	test.insertToTrie("21e8", "21eabf80faebc12002aec48f82ba433758130924fde0c0b03dace7b0c9c42f09")

	test.insertToTrie("21e", "21e813411aa5c7f3896fe57fb7d8a8f150ee18a7256fe73990a17c47a498c8b5")
	//End Test.

	// InitSocket()
	r := mux.NewRouter()

	r.HandleFunc("/api/v2/mine", MineReq).Methods("POST")
	r.HandleFunc("/api/v2/hashwall", Hashwall).Methods("PUT")
	r.HandleFunc("/api/v2/data/{dataHash}", GetData).Methods("GET")
	r.HandleFunc("/api/v2/index/{sourceHash}", GetIndex).Methods("GET")
	r.HandleFunc("/api/v2/trie/{target}", TriePrefixLookup).Methods("GET")
	r.HandleFunc("/api/v2/raw/{rotation}", GetRaw).Methods("GET")
	// http.HandleFunc("/", wsEndpoint)

	fmt.Println("Running node on port 3222.")
	// go func() {
	// 	http.ListenAndServe(":2180", nil)
	// 	fmt.Println("WS Server Started on port 2180")
	// }()

	http.ListenAndServe(":3222", r)
}
