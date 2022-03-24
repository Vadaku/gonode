package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"time"
)

type MineResult struct {
	source   string
	datahash string
	target   string
	rotation string
	nonce    string
}

//Mining function to handle mine endpoint.
//Processes clients mine request then returns mine result as response.
func mine(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved mine request from", r.RemoteAddr)
	//Assign request queries and random nonce to relevant variables.
	r.ParseForm()
	source := r.FormValue("source")
	data := r.FormValue("data")
	target := r.FormValue("target")
	rand.Seed(time.Now().UnixNano())
	nonce := rand.Int()
	//Hash the source if it hasn't been hashed already.
	if len(source) != 64 {
		sourceChecksum := sha256.Sum256([]byte(source))
		source = hex.EncodeToString(sourceChecksum[:])
	}
	//Hash the data if it hasn't been hashed already. --> Passed datahash because of hashwall?
	if len(data) != 64 {
		dataChecksum := sha256.Sum256([]byte(data))
		data = hex.EncodeToString(dataChecksum[:])
	}
	//Hash data variable.
	dataChecksum := sha256.Sum256([]byte(data))
	dataHash := hex.EncodeToString(dataChecksum[:])
	//Hash combined "rotation" string.
	rotationChecksum := sha256.Sum256([]byte(source + dataHash + strconv.Itoa(nonce)))
	rotationHash := hex.EncodeToString(rotationChecksum[:])
	//Continually mine to see if the rotationHash prefix matches the target.
	for {
		if rotationHash[0:len(target)] == target {
			fmt.Printf("\033[32mTarget matched with rotation %s and nonce %s\033[0m\n", rotationHash, strconv.Itoa(nonce))
			break
		}
		nonce++
		rotationChecksum = sha256.Sum256([]byte(source + dataHash + strconv.Itoa(nonce)))
		rotationHash = hex.EncodeToString(rotationChecksum[:])
	}

	//Format result as MineResult struct.
	_ = &MineResult{
		source:   source,
		datahash: dataHash,
		target:   target,
		rotation: rotationHash,
		nonce:    strconv.Itoa(nonce),
	}

	//Add data to leveldb.
	AddToData(dataHash, data)
}

//Return index data to client
func getIndex(w http.ResponseWriter, r *http.Request) {

}

//Mine a rotation to unlock data given the datahash.
//Expects source, datahash, target and rotation as inputs.
func hashwall(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// source := r.FormValue("source")
	// datahash := r.FormValue("datahash")
	// target := r.FormValue("target")
}

//Return data to the client
func getData(w http.ResponseWriter, r *http.Request) {
	dataHash := path.Base(r.URL.Path)
	fmt.Printf("Recieved data request for %s\n", dataHash)
	result := DBGetData(dataHash)
	w.Write([]byte(result))
}
