package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type MineResult struct {
	source   string
	datahash string
	target   string
}

//Mining function to handle mine endpoint.
//Processes clients mine request then returns mine result as response.
func mine(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved mine request from ", r.RemoteAddr)
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
	//Hash data variable.
	dataChecksum := sha256.Sum256([]byte(data))
	dataHash := hex.EncodeToString(dataChecksum[:])
	//Hash combined "rotation" string.
	rotationChecksum := sha256.Sum256([]byte(source + dataHash + strconv.Itoa(nonce)))
	rotationHash := hex.EncodeToString(rotationChecksum[:])
	//Continually check to see if the rotationHash prefix matches the target.
	//If not, increment the nonce by 1 then hash the new rotation string and continue.
	//If the prefix matches then execute the loop and return result.
	for {
		if rotationHash[0:len(target)] == target {
			fmt.Printf("Target matched with rotation %s and nonce %s\n", rotationHash, strconv.Itoa(nonce))
			break
		}
		nonce++
		rotationChecksum = sha256.Sum256([]byte(source + dataHash + strconv.Itoa(nonce)))
		rotationHash = hex.EncodeToString(rotationChecksum[:])
	}
	//Add data to leveldb.
	AddToData(dataHash, data)
}

//Return index data to client
func getIndex(w http.ResponseWriter, r *http.Request) {

}

//Mine a rotation to unlock data given the datahash
func hashwall(w http.ResponseWriter, r *http.Request) {

}

//Return data to the client
func getData(w http.ResponseWriter, r *http.Request) {

}
