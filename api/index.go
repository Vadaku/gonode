package api

import (
	"gonode/database"
	"log"
	"net/http"
	"path"

	"github.com/syndtr/goleveldb/leveldb"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	root := path.Base(r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "GET" {
		http.Error(w, "Expected GET request", http.StatusNotFound)
	} else {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		db, err := leveldb.OpenFile("../.history/index", nil)
		if err != nil {
			log.Fatal(err)
		}
		x, y := database.DBIndexLookup(root, *db)
		if y {
			println("Sending to client")
		}
		w.Write(x)
		db.Close()
	}
}
