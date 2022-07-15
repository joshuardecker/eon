package main

import (
	"bytes"
	"fmt"

	"github.com/GoblinBear/beson-go/types"
	"golang.org/x/crypto/sha3"
)

// Used for easy conversion of uint32 and []byte
func i32tob(val uint32) []byte { // TODO: make not copywritted
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}

func main() {

	//****
	// Section Uncompacts the target

	// Target compact form: 0x1d00ffff, this compact form is now unpacked into a 256 bit uint form
	compactTarget := uint32(0x1d00000f)               // Defines the compact uint
	compactTargetExponent := compactTarget >> (3 * 8) // Gets the first byte out of it (big endian style)

	mantissa := types.NewUInt256("0", 1)                               // Defines a blank canvis that is a 256 bit uint
	mantissa.Set(i32tob(compactTarget & 0x00ffffff))                   // Masks the first byte and inputs the rest into the uint235
	mantissa = mantissa.LShift(uint(8 * (32 - compactTargetExponent))) // To deal with the little endian that is this, shifts it left(right in big endian) to properly format

	target := types.NewUInt256("0", 1) // TODO: shorten if possible
	target.Set(mantissa.Get())

	// Section Uncompacts the target
	//****

	//****
	// Starts Mining
	for nonce := uint32(0); nonce <= 0xFFFFFFFF; nonce += 1 {
		data := []byte("Hello World!")
		data = append(data[:], i32tob(nonce)...) //TODO: look up the ... opertator // Appends the nonce and data, nonce needs to be added seperate from the data cause funny

		h := make([]byte, 32) // Prepares the hash var

		sha3.ShakeSum256(h, data) // Hashes the custom sized digest 'h'

		// Is the hash less or = to the target
		if bytes.Compare(h, target.Get()) != 1 {
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
			fmt.Printf("Target: %x\n", target.Get())
		}
	}
	// Starts the mining
	//****
}
