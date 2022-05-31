package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func Mine(source string, data string, target string, conn *websocket.Conn) *MineResult {
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
		fmt.Println(rotationHash)
		if conn != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(rotationHash))
		}
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
		Weight:    int64(math.Pow(16, float64(len(target)))),
	}
	fmt.Println(timestamp)
	weight := strconv.FormatInt(int64(math.Pow(16, float64(len(target)))), 10)
	//Hardcoded 'user' hash.
	myNameHash := "00e51906df651a7ee922446590f487cff433ec9816aedc44dc49952a05cd16df"
	rawString := strconv.FormatInt(timestamp, 10) + weight + source + data + target + myNameHash + strconv.Itoa(nonce)

	//Add data to leveldb.

	AddToData(data, beforeHashData)
	AddToIndex(source, rotationHash, rawString)

	return res
}

func BinaryData() {
	fmt.Println("In BinaryData()")
}
