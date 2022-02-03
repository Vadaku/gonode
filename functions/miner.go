package functions

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gonode/database"
	"gonode/result"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

func MinerFork(source string, data string, target string) {
	nonce := 1
	if len(source) != 64 {
		hashingSource := sha256.Sum256([]byte(source))
		source = hex.EncodeToString(hashingSource[:])
	}
	hashingData := sha256.Sum256([]byte(data))
	hashedData := hex.EncodeToString(hashingData[:])
	// if err != nil {
	// 	panic(err)
	// }
	rotation := source + hashedData + strconv.Itoa(nonce)
	rotationHash := sha256.Sum256([]byte(rotation))
	rotationHashed := hex.EncodeToString(rotationHash[:])
	for {
		if rotationHashed[0:len(target)] == target {
			fmt.Println("Target matched", rotationHashed, "with nonce", nonce)
			break
		}
		nonce++
		rotation = source + hashedData + strconv.Itoa(nonce)
		rotationHash = sha256.Sum256([]byte(rotation))
		rotationHashed = hex.EncodeToString(rotationHash[:])
	}
	res := &result.Result{
		Datahash:  hashedData,
		Nonce:     strconv.Itoa(nonce),
		Rotation:  rotationHashed,
		Source:    source,
		Target:    target,
		Timestamp: time.Now().String(),
		User:      "Anonynmous",
	}

	output, err := proto.Marshal(res)
	if err != nil {
		panic(err)
	}
	//Add mining result to DB.
	database.DBAddResult(rotationHashed, output)
	//Add datahash as "key" and data as "value" to DB.
	database.DBAddData(hashedData, data)
}

func MinerRotate(source string, data string, target string) {

}