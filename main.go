package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var test *Trie

//Setup routes and handlers then serve on port 8080.
func main() {
	//Testing Trie.
	test = initializeTrie()
	test.insertToTrie("21e8", "21e893411ac5c7f3896fe57fb7d8a8f150ee18a7256fe73990a17c47a498c8b5")
	test.insertToTrie("21e8", "testrotation")
	test.insertToTrie("21e8", "testrotation123")

	test.insertToTrie("21e", "21e893411ac5c7f3896fe57fb7d8a8f150ee18a7256fe73990a17c47a498c8b5")
	fmt.Println(test.searchTrie("21e8"))
	fmt.Println(test.rootNode)
	//End Test.

	r := mux.NewRouter()

	r.HandleFunc("/api/v2/mine", mine).Methods("POST")
	r.HandleFunc("/api/v2/hashwall", hashwall).Methods("PUT")
	r.HandleFunc("/api/v2/data/{dataHash}", getData).Methods("GET")
	r.HandleFunc("/api/v2/index/{sourceHash}", getIndex).Methods("GET")
	r.HandleFunc("/api/v2/trie/{target}", triePrefixLookup).Methods("GET")

	fmt.Println("Running node on port 3222.")
	http.ListenAndServe(":3222", r)
}
