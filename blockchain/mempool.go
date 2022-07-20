package blockchain

import "github.com/Sucks-To-Suck/LuncheonNetwork/transactions"

type Mempool struct {
	txs []transactions.LuTx
}

// Function adds a tx to the mempool of the blockchain.
// Inputs the transaction you are adding.
// Returns nothing.
func (m *Mempool) AddTx(tx transactions.LuTx) {

	m.txs = append(m.txs, tx)
}
