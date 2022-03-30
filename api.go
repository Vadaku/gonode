package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"time"
)

type MineResult struct {
	Source    string `json:"source"`
	Datahash  string `json:"datahash"`
	Target    string `json:"target"`
	Rotation  string `json:"rotation"`
	Nonce     string `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
}

//Mining function to handle mine endpoint.
//Processes clients mine request then returns mine result as response.
func MineReq(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved mine request from", r.RemoteAddr)
	r.ParseForm()
	source := r.FormValue("source")
	data := r.FormValue("data")
	target := r.FormValue("target")
	if source == "" || data == "" || target == "" {
		http.Error(w, "Missing a required parameter.\nPlease ensure request includes source, data and target.", http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Mine(source, data, target))
	}

}

func Mine(source string, data string, target string) *MineResult {
	var beforeHashData string
	//Assign request queries and random nonce to relevant variables.
	rand.Seed(time.Now().UnixNano())
	nonce := rand.Int()
	//Hash the source if it hasn't been hashed already.
	if len(source) != 64 {
		sourceChecksum := sha256.Sum256([]byte(source))
		source = hex.EncodeToString(sourceChecksum[:])
	}
	//Hash the data if it hasn't been hashed already. --> Passed datahash because of hashwall?
	if len(data) != 64 {
		beforeHashData = data
		dataChecksum := sha256.Sum256([]byte(data))
		data = hex.EncodeToString(dataChecksum[:])
	}
	//Hash combined "rotation" string.
	rotationChecksum := sha256.Sum256([]byte(source + data + strconv.Itoa(nonce)))
	rotationHash := hex.EncodeToString(rotationChecksum[:])
	//Continually mine to see if the rotationHash prefix matches the target.
	for {
		if rotationHash[0:len(target)] == target {
			fmt.Printf("\033[32mTarget matched with rotation %s and nonce %s\033[0m\n", rotationHash, strconv.Itoa(nonce))
			break
		}
		nonce++
		rotationChecksum = sha256.Sum256([]byte(source + data + strconv.Itoa(nonce)))
		rotationHash = hex.EncodeToString(rotationChecksum[:])
	}

	timestamp := time.Now().Unix()
	//Format result as MineResult struct.
	res := &MineResult{
		Source:    source,
		Datahash:  data,
		Target:    target,
		Rotation:  rotationHash,
		Nonce:     strconv.Itoa(nonce),
		Timestamp: timestamp,
	}
	rawString := source + data + target + rotationHash + strconv.Itoa(nonce)

	//Add data to leveldb.

	AddToData(data, beforeHashData)
	AddToIndex(source, rotationHash, rawString)

	return res
}

//Return index data to client
func GetIndex(w http.ResponseWriter, r *http.Request) {
	source := path.Base(r.URL.Path)
	res, _ := DBGetIndex(source)
	w.Write([]byte(res))

}

//Mine a rotation to unlock data given the datahash.
//Expects mine result as inputs.
func Hashwall(w http.ResponseWriter, r *http.Request) {
	var reqBody MineResult
	header := r.Header.Get("Content-Type")
	if header != "" {
		if header != "application/json" {
			http.Error(w, "Expected application/json Content-Type header", http.StatusUnsupportedMediaType)
		} else {
			err := json.NewDecoder(r.Body).Decode(&reqBody)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				Mine(reqBody.Source, reqBody.Datahash, reqBody.Target)
				http.Redirect(w, r, "/api/v2/data/"+reqBody.Datahash, http.StatusMovedPermanently)
			}
		}
	}
}

//Return data to the client
func GetData(w http.ResponseWriter, r *http.Request) {
	dataHash := path.Base(r.URL.Path)
	fmt.Printf("Recieved data request for %s\n", dataHash)
	result := DBGetData(dataHash)
	w.Write([]byte(result))
}

func GetRaw(w http.ResponseWriter, r *http.Request) {
	rotation := path.Base(r.URL.Path)
	fmt.Printf("Recieved request for raw meta.\n")
	result, _ := DBGetIndex(rotation)
	w.Write([]byte(result))
}

//Trie lookup using a target then return rotations associated with the target.
func TriePrefixLookup(w http.ResponseWriter, r *http.Request) {
	target := path.Base(r.URL.Path)
	result, _ := test.searchTrie(target)
	fmt.Println(result)
}
