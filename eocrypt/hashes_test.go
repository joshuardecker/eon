package eocrypt

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/Sucks-To-Suck/Eon/tools/logger"
)

func TestHashes(t *testing.T) {

	// Test value setting:

	h := new(Hash)

	// Test string.
	h.SetString("This is a test")
	logger.LogYellow("Hash Tester", h.String())

	// Test hex.
	h.SetHex("A13891BFFCD")
	logger.LogYellow("Hash Tester", h.Hex())

	// Test bytes.
	b := []byte("This is a test!")
	fmt.Printf("Bytes: %x\n", b)
	h.SetBytes(b)
	fmt.Printf("Bytes: %x\n", h.Bytes())

	// Test Big Int.
	// This will return a number much higher than 1, but that is alright.
	// This is because anything inputted into a Big Int is inputted left to right, not right to left.
	i := big.NewInt(1)
	h.SetBigInt(i)
	logger.LogYellow("Hash Tester", h.BigInt().String())

	// Test actually hashing:

	// Hash some bytes.
	fmt.Printf("Hash: %x\n", *HashBytes(b))

	// Hash an interface of any data!
	fmt.Printf("Hash: %x\n", *HashInterface(
		[]interface{}{
			8992,
			"This is a test",
			"Sample data",
			7.98919,
		}))
}
