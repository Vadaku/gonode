package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

type MineResult struct {
	Source    string `json:"source"`
	Datahash  string `json:"datahash"`
	Target    string `json:"target"`
	Rotation  string `json:"rotation"`
	Nonce     string `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
	Weight    int64  `json:"weight"`
}

type RotationList struct {
	Rotations []string `json:"rotations"`
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
		jsonResult, _ := Mine(source, data, target, nil)
		json.NewEncoder(w).Encode(jsonResult)
	}

}

//Return index data to client
func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rotations RotationList
	source := path.Base(r.URL.Path)
	res, _ := DBGetIndex(source)
	rotations.Rotations = res
	json.NewEncoder(w).Encode(rotations)
	// w.Write([]byte(res))
}

//Mine a rotation to unlock data given the datahash.
//Expects mine result as inputs.
func Hashwall(w http.ResponseWriter, r *http.Request) {
	// ws, _ := upgrader.Upgrade(w, r, nil)
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
				Mine(reqBody.Source, reqBody.Datahash, reqBody.Target, nil)
				http.Redirect(w, r, "/api/v2/data/"+reqBody.Datahash, http.StatusMovedPermanently)
			}
		}
	}
}

//Return data to the client
func GetData(w http.ResponseWriter, r *http.Request) {
	dataHash := path.Base(r.URL.Path)
	fmt.Printf("Recieved data request for %s\n", dataHash)
	header := r.Header.Get("Content-Type")
	if header != "" {
		if header != "application/octet-stream" && header != "text/plain" {
			http.Error(w, "Expected application/octect or text/plain", http.StatusUnsupportedMediaType)
		} else if header == "text/plain" {

		} else if header == "application/octet-stream" {
			BinaryData()
		}
	}
	result := DBGetData(dataHash)
	w.Write([]byte(result))
}

func GetRaw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rotation := path.Base(r.URL.Path)
	fmt.Printf("Recieved request for raw meta.\n")
	result, _ := DBGetIndex(rotation)
	meta := map[string]interface{}{
		rotation: strings.Join(result, ""),
	}
	json.NewEncoder(w).Encode(meta)
	// w.Write([]byte(result))
}

func GetJson(w http.ResponseWriter, r *http.Request) {

}

//Trie lookup using a target then return rotations associated with the target.
func TriePrefixLookup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rotations RotationList
	target := path.Base(r.URL.Path)
	result, _ := test.searchTrie(target)
	rotations.Rotations = result
	json.NewEncoder(w).Encode(rotations)
}

func PostBinary(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	tmpfile, _ := os.Create("./" + "photo.png")
	defer tmpfile.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	tmpfile.Write(d)
	w.WriteHeader(200)
}
