package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func AddToData(dataHash string, data string) {
	db, _ := leveldb.OpenFile("../.history/data", nil)

	if data != "" {
		if ok, _ := db.Has([]byte(dataHash), nil); !ok {
			db.Put([]byte(dataHash), []byte(data), nil)
		}
	}

	defer db.Close()
}

func AddToIndex(source string, result string, raw string) {
	db, _ := leveldb.OpenFile("../.history/index", nil)
	rotation := result
	indexResult, inDb := db.Get([]byte(source), nil)
	if inDb == nil {
		result += string(indexResult)
	}
	db.Put([]byte(source), []byte(result), nil)
	db.Put([]byte(rotation), []byte(raw), nil)

	AddRaw(result, raw, db)
	defer db.Close()
}

func AddRaw(rotation string, rawstring string, db *leveldb.DB) {
}

func DBGetIndex(key string) (string, bool) {
	inDb := true
	db, _ := leveldb.OpenFile("../.history/index", nil)
	res, _ := db.Get([]byte(key), nil)
	// if ok != nil {
	// 	inDb = false
	// }

	defer db.Close()
	return string(res), inDb
}

func DBGetData(dataHash string) string {
	db, _ := leveldb.OpenFile("../.history/data", nil)

	data, err := db.Get([]byte(dataHash), nil)

	if err != nil {
		fmt.Printf("Database does not contain key %s\n", dataHash)
	}

	defer db.Close()
	return string(data)
}
