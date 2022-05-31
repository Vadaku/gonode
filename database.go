package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

func AddToTrie() {
	test = initializeTrie()

	f, err := os.Open("../.history/index/")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		test.insertToTrie("21e8", v.Name()[0:64])
	}
}

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
	rotationFile, _ := os.Create("../.history/index/" + result + ".txt")
	sourceFile, err := os.OpenFile("../.history/index/"+source+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer sourceFile.Close()
	defer rotationFile.Close()

	_, err2 := rotationFile.WriteString(raw)
	_, _ = sourceFile.WriteString(result + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}
}

func AddRaw(rotation string, rawstring string, db *leveldb.DB) {
}

func DBGetIndex(key string) ([]string, bool) {
	inDb := true
	sourceFile, err := os.Open("../.history/index/" + key + ".txt")

	var words []string

	if err != nil {
		log.Fatal(err)
	}

	Scanner := bufio.NewScanner(sourceFile)
	Scanner.Split(bufio.ScanWords)

	for Scanner.Scan() {
		words = append(words, Scanner.Text())
	}
	if err := Scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer sourceFile.Close()

	return words, inDb
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
