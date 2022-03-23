package main

import (
	"fmt"

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

func DBGetData(dataHash string) string {
	db, _ := leveldb.OpenFile("../.history/data", nil)

	data, err := db.Get([]byte(dataHash), nil)

	if err != nil {
		fmt.Printf("Database does not contain key %s\n", dataHash)
	}
	return string(data)
}
