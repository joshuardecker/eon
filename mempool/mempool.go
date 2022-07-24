package mempool

import (
	"encoding/hex"

	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
	"github.com/Sucks-To-Suck/LuncheonNetwork/wallet"
	"golang.org/x/crypto/sha3"
)

// The mempool struct, containing all the tx's waiting to be added to the next available block.
type Mempool struct {
	Txs []transactions.LuTx

	wal *wallet.Wallet
}

// Initialize the mempool with a wallet.
// Returns the initialized mempool.
func Init(wal *wallet.Wallet) Mempool {

	m := new(Mempool)

	m.wal = wal

	return *m
}

// Function adds a tx to the mempool of the blockchain.
// Inputs the tx you are adding.
// Returns true if successfully added, false if tx was invalid.
func (m *Mempool) AddTx(tx transactions.LuTx) bool {

	// If the tx has a spendable amount of coin from the persons balance
	if (tx.Value + tx.Fee) > m.wal.ScanChainForBalance(tx.TxFrom) {

		return false
	}

	// If the tx has the wrong nonce value
	if tx.Nonce != m.wal.ScanChainForNonce(tx.TxFrom) {

		return false
	}

	// Remove the sig from the tx and save it, as to get the tx hash input data
	signature, _ := hex.DecodeString(tx.Signature)
	tx.Signature = ""

	txBytes := tx.AsBytes()
	txHash := make([]byte, 32)
	pubKey, _ := hex.DecodeString(tx.TxFrom)

	sha3.ShakeSum256(txHash, txBytes)

	// If the signature is not valid
	if !ellip.ValidateSig(pubKey, txHash, signature) {

		return false
	}

	m.Txs = append(m.Txs, tx)
	return true
}

// This function removes a tx from the mempool.
// Returns nothing.
func (m *Mempool) RemoveTx(index int) {

	m.Txs = append(m.Txs[:index], m.Txs[index+1:]...)
}

// Gets and returns a valid tx.
func (m *Mempool) GetTx() transactions.LuTx {

	// If no txs
	if len(m.Txs) == 0 {

		return transactions.LuTx{}
	}

	tx := m.Txs[0]
	m.RemoveTx(0)

	return tx
}
