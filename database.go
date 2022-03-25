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

func AddToIndex(source string, result string) {
	db, _ := leveldb.OpenFile("../.history/index", nil)

	indexResult, inDb := db.Get([]byte(source), nil)
	if inDb == nil {
		result += string(indexResult)
	}

	fmt.Println("Incoming result", result)

	db.Put([]byte(source), []byte(result), nil)

	defer db.Close()
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
