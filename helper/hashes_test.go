package helper

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/Sucks-To-Suck/LuncheonNetwork/helper/logger"
)

func TestHashes(t *testing.T) {

	// Test value setting:

	logger := logger.NewLogger("Test")

	h := EmptyHash()

	// Test string.
	h.SetString("This is a test")
	logger.LogYellow(h.GetString())

	// Test hex.
	h.SetHex("A13891BFFCD")
	logger.LogYellow(h.GetHex())

	// Test bytes.
	b := []byte("This is a test!")
	fmt.Printf("Bytes: %x\n", b)
	h.SetBytes(b)
	fmt.Printf("Bytes: %x\n", h.GetBytes())

	// Test Big Int.
	// This will return a number much higher than 1, but that is alright.
	// This is because anything inputted into a Big Int is inputted left to right, not right to left.
	i := big.NewInt(1)
	h.SetBigInt(i)
	logger.LogYellow(h.GetBigInt().String())

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
