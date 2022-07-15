package main

import (
	"bytes"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/luncheon"
	"golang.org/x/crypto/sha3"
)

func main() {

	//****
	// Section Unpacks the target

	unpacker := new(luncheon.TargetUnpacker)
	target := unpacker.UnpackAsBytes(0x1d00ffff)

	// Section Unpacks the target
	//****

	//****
	// Starts Mining
	for nonce := uint32(0); nonce <= 0xFFFFFFFF; nonce += 1 {
		data := []byte("Hello World")
		data = append(data[:], luncheon.Uint32toB(nonce)...) //TODO: look up the ... opertator // Appends the nonce and data, nonce needs to be added seperate from the data cause funny

		h := make([]byte, 32) // Prepares the hash var

		sha3.ShakeSum256(h, data) // Hashes the custom sized digest 'h'

		// Is the hash less or = to the target
		if bytes.Compare(h, target) != 1 {
			fmt.Println("Hash is valid!")
			fmt.Printf("Hash: %x\n", h)
			fmt.Println("Nonce: ", nonce)

			break
		}

		// Prints stats every 10 MH
		if nonce%10000000 == 0 {
			fmt.Println("Hashing...")
			fmt.Printf("Hash: %x\n", h)
			fmt.Println("Nonce: ", nonce)
			fmt.Printf("Target: %x\n", target)
		}
	}
	// Starts the mining
	//****
}
