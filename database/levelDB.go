package database

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func DBAddResult(rotation string, r []byte) {
	db, err := leveldb.OpenFile("../.history/results", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Adding %s result to database.\n", rotation)
	db.Put([]byte(rotation), r, nil)

	db.Close()
}

func DBAddData(datahash string, data string) {
	db, err := leveldb.OpenFile("../.history/data", nil)
	if err != nil {
		panic(err)
	}
	db.Put([]byte(datahash), []byte(data), nil)

	db.Close()
}

func DBDataLookup(datahash string) []byte {
	db, err := leveldb.OpenFile("../.history/data", nil)
	if err != nil {
		panic(err)
	}
	data, err := db.Get([]byte(datahash), nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data being sent to client\n")
	db.Close()
	return data
}

func DBIndexLookup(root string) {

}
