package eocrypt

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"math/big"

	"github.com/Sucks-To-Suck/Eon/tools/logger"
	"golang.org/x/crypto/sha3"
)

const HASH_LEN = 32

// Is only a 32 long byte array, but with some fancy functions.
type Hash [HASH_LEN]byte

// ****
// Hash struct funcs:

// Set the value of the hash with bytes.
func (h *Hash) SetBytes(b []byte) {

	if len(b) > 32 {

		// If the bytes are longer than 32, only copy the first 32 bytes (from right -> left in the byte array).
		b = b[len(b)-HASH_LEN:]
	}

	copy(h[:], b)
}

// Set the value of the hash with a string.
func (h *Hash) SetString(s string) {

	h.SetBytes([]byte(s))
}

// Set the value of the hash with a hex string.
func (h *Hash) SetHex(s string) {

	// Decode the hex.
	b, _ := hex.DecodeString(s)

	h.SetBytes(b)
}

// Set the value of the hash with a Big Int.
func (h *Hash) SetBigInt(i *big.Int) {

	h.SetBytes(i.Bytes())
}

// Get the bytes of the hash.
func (h *Hash) Bytes() []byte {

	return h[:]
}

// Get the string of the hash.
func (h *Hash) String() string {

	return string(h[:])
}

// Get the string (hex) of the bytes of the hash.
func (h *Hash) Hex() string {

	return hex.EncodeToString(h[:])
}

// Get the Big Int of the bytes of the hash.
func (h *Hash) BigInt() *big.Int {

	// Create and assign a value to the big int.
	i := big.NewInt(0)
	i.SetBytes(h[:])

	return i
}

// Hash struct funcs:
// ****

// ****
// Hashing Processes:

// Hash input bytes and return it as a hash type.
func HashBytes(b []byte) *Hash {

	h := new(Hash)

	// Hash using shake.
	sha3.ShakeSum128(h[:], b)

	return h
}

// This function hashes an inputted interface.
// Returns nil and logs if an error occurs.
func HashInterface(val interface{}) *Hash {

	h := new(Hash)

	// Make a buffer.
	buf := new(bytes.Buffer)

	// Make an encoder and encode the interface 'val'.
	err := gob.NewEncoder(buf).Encode(val)

	// If an error occured, log it in red and return nil.
	if err != nil {

		logger.LogRed("Hasher", "Couldnt hash interface! "+err.Error())

		return nil
	}

	// Hash the bytes.
	sha3.ShakeSum128(h[:], buf.Bytes())

	return h
}

// Mixes two given hashes together, in a very lightweight process.
func MixHashes(h1 *Hash, h2 *Hash) *Hash {

	// Make a byte array the size of the hash input.
	hBytes := make([]byte, len(h1.Bytes()))

	// Loop through all of the bytes and mix (XOR) them together.
	for i := range h1 {

		hBytes[i] = h1[i] ^ h2[i]
	}

	// Return it as a hash.
	return HashBytes(hBytes)
}

// Hashing Processes:
// ****
