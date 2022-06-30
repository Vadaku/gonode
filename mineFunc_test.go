package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var s string
var d string
var t string
var nonce int

func TestMain(m *testing.M) {
	flag.StringVar(&s, "s", "", "Source hash")
	flag.StringVar(&d, "d", "", "Data hash")
	flag.StringVar(&t, "t", "", "Target prefix")
	flag.IntVar(&nonce, "nonce", 500, "Nonce")
	flag.Parse()

	os.Exit(m.Run())

}

func TestMineFunction(testing *testing.T) {
	correctMineRes := &MineResult{}

	flagNonce := nonce

	rotationTest := mineFunction(s, d, t, flagNonce)

	correctMineRes, _ = Mine(s, d, t, flagNonce, false)
	fmt.Printf("Test: %s\nCorrect: %s\n", rotationTest, correctMineRes.Rotation)

	if correctMineRes.Rotation != rotationTest {
		testing.Errorf("Rotation not correct.")
	}
}
