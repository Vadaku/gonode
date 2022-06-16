package main
	import (
		_"crypto/sha256"
		_"encoding/hex"
		"strconv"
	)/**
	Source and data parameters are sha256 hashes.
	Target is a valid target.
**/
func mineFunction(source string, data string, target string, nonce int) string {
	rotationHash := source + data + strconv.Itoa(nonce)
	_ = rotationHash
	//Complete code below...

	return ""
}
