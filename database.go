package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

func AddToData(dataHash string, data string) {
	db, _ := leveldb.OpenFile("../.history/data", nil)

	db.Put([]byte(dataHash), []byte(data), nil)

	defer db.Close()
}

func AddToIndex(result string) {
	db, _ := leveldb.OpenFile("../.history/index", nil)

	defer db.Close()
}
