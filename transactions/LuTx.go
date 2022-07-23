package transactions

import (
	"encoding/hex"
	"encoding/json"

	"golang.org/x/crypto/sha3"
)

// This struct are the tx's on the Luncheon Network.
type LuTx struct {
	txFrom string
	txTo   string
	value  uint64

	script string

	nonce     uint32
	fee       uint64
	signature string
}

// Sets the script of the tx.
// If a blank script is entered, nothing is inputted into the tx.
// Returns nothing.
func (l *LuTx) AddScriptStr(scriptstr string) {

	if len(scriptstr) == 0 {

		return
	}

	// Runs the tx through the scripter and back out to remove any junk
	l.script = ScriptToStr(StrToScript(scriptstr))
}

// Function converts the tx into bytes.
// Returns the byte array of the tx.
func (l *LuTx) AsBytes() []byte {

	lAsBytes, jsonErr := json.Marshal(l)

	if jsonErr != nil {

		panic(jsonErr)
	}

	return lAsBytes
}

// This function calculates the hash of the transaction.
// Returns the string hex of the transaction hash.
func (l *LuTx) HashTx() string {

	hash := make([]byte, 32)
	sha3.ShakeSum256(hash, l.AsBytes())

	return hex.EncodeToString(hash)
}
