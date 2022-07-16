package main

import (
	"bytes"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/luncheon"
	"golang.org/x/crypto/sha3"
)

func main() {

	//****
	// Init the needed structs

	util := new(luncheon.Util)
	unpacker := new(luncheon.TargetUnpacker)
	packer := new(luncheon.TargetPacker)

	// Init the needed structs
	//****

	//****
	// Section Unpacks the target

	target := unpacker.UnpackAsBytes(0x1d00ffff)

	// Recalculates the packed target, right now thats for debugging purposes / testing
	packedTarget, packErr := packer.PackTargetBytes(target)

	// Did that func through errors?
	if packErr != nil {

		panic(packErr)
	}

	fmt.Printf("Packed Target: %x\n", 0x1d00ffff)
	fmt.Printf("Recalculated Packed Target: %x\n", packedTarget)

	// Section Unpacks the target
	//****

	//****
	// Starts Mining

	// Init the nonce value, starting with 0
	for nonce := uint32(0); nonce <= 0xFFFFFFFF; nonce += 1 {

		// Random data used in place of block data for now
		data := []byte("Hello World")

		// Combines the data and a non size limited nonce
		data = append(data[:], util.Uint32toB(nonce)...)

		// Prepares the hash (32 bytes)
		h := make([]byte, 32)

		// Hashes the data into h, h can be any custom size too
		sha3.ShakeSum256(h, data)

		// Is the hash less or = to the target
		if bytes.Compare(h, target) != 1 {

			fmt.Println("Hash is valid!")
			fmt.Printf("Hash: %x\n", h)
			fmt.Println("Nonce: ", nonce)
			fmt.Printf("Target: %x\n", target)

			break
		}

		// Prints stats every 10 MH's
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
