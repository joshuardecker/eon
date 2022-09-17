package helper

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"math/big"

	"github.com/Sucks-To-Suck/LuncheonNetwork/helper/logger"
	"golang.org/x/crypto/sha3"
)

const HASH_LEN = 32

// Is only a 32 long byte array, but with some fancy functions.
type hash [HASH_LEN]byte

// ****
// Value setting / getting:

// Set the value of the hash with bytes.
func (h *hash) SetBytes(b []byte) {

	if len(b) > 32 {

		// If the bytes are longer than 32, only copy the first 32 bytes
		b = b[len(b)-HASH_LEN:]
	}

	copy(h[:], b)
}

// Set the value of the hash with a string.
func (h *hash) SetString(s string) {

	h.SetBytes([]byte(s))
}

// Set the value of the hash with a hex string.
func (h *hash) SetHex(s string) {

	// Decode the hex.
	b, _ := hex.DecodeString(s)

	h.SetBytes(b)
}

// Set the value of the hash with a Big Int.
func (h *hash) SetBigInt(i *big.Int) {

	h.SetBytes(i.Bytes())
}

// Get the bytes of the hash.
func (h *hash) GetBytes() []byte {

	return h[:]
}

// Get the string of the hash.
func (h *hash) GetString() string {

	return string(h[:])
}

// Get the string (hex) of the bytes of the hash.
func (h *hash) GetHex() string {

	return hex.EncodeToString(h[:])
}

// Get the Big Int of the bytes of the hash.
func (h *hash) GetBigInt() *big.Int {

	// Create and assign a value to the big int.
	i := big.NewInt(0)
	i.SetBytes(h[:])

	return i
}

// Value setting / getting:
// ****

// ****
// Usage:

// Create and return an emtpy hash.
func EmptyHash() *hash {

	return new(hash)
}

// Hash input bytes and return it as a hash type.
func HashBytes(b []byte) *hash {

	h := new(hash)

	// Hash using shake.
	sha3.ShakeSum128(h[:], b)

	return h
}

// This function hashes an inputted interface.
// Returns nil and logs if an error occurs.
func HashInterface(val interface{}) *hash {

	h := new(hash)

	// Make a buffer.
	var buf bytes.Buffer

	// Make an encoder and encode the interface 'i'.
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(val)

	// If an error occured, log it in red and return nil.
	if err != nil {

		logger := logger.NewLogger("Helper")

		logger.LogRed("Couldnt hash interface! " + err.Error())

		return nil
	}

	// Hash the bytes.
	sha3.ShakeSum128(h[:], buf.Bytes())

	return h
}

// Usage:
// ****
