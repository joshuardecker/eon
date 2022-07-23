package blockchain

import (
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

// The mempool struct, containing all the tx's waiting to be added to the next available block.
type Mempool struct {
	Txs []transactions.LuTx
}

// Function adds a tx to the mempool of the blockchain.
// Inputs the transaction you are adding.
// Returns nothing.
func (m *Mempool) AddTx(tx transactions.LuTx) {

	m.Txs = append(m.Txs, tx)
}

// This function removes a tx from the mempool.
// Returns nothing.
func (m *Mempool) RemoveTx(index int) {

	m.Txs = append(m.Txs[:index], m.Txs[index+1:]...)
}
