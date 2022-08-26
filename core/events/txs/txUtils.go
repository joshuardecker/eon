package txs

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"

	"github.com/Sucks-To-Suck/LuncheonNetwork/crypto/ellip"
)

func CreateTx(fromPub string, toPub string, amount uint64, nonce uint32,
	fromPriv *ecdsa.PrivateKey) (tx LuTx) {

	// Fill in the tx
	tx.TxFrom = fromPub
	tx.TxTo = toPub
	tx.Value = amount
	tx.Nonce = nonce

	// Simple calculation to get a tx fee
	tx.Fee = uint64((tx.GetWeight() + 64) * 100) // The +64 is to add the weight of the signature

	_, sig := ellip.SignMsg(fromPriv, tx.AsBytes())
	tx.Signature = hex.EncodeToString(sig)

	return tx
}

// Gets the tx as bytes.
// Returns nil if an error occured.
func (t *LuTx) TxBytes() []byte {

	// Get the bytes
	bytes, err := json.Marshal(t)

	if err != nil {

		return nil
	}

	return bytes
}
