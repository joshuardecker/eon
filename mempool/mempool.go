package blockchain

import (
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
	"github.com/Sucks-To-Suck/LuncheonNetwork/wallet"
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
// Inputs the transaction you are adding.
// Returns nothing.
func (m *Mempool) AddTx(tx transactions.LuTx) {

	// If the tx has a spendable amount of coin from the persons balance
	if m.wal.CheckTxAmount(tx) {

		m.Txs = append(m.Txs, tx)
	}
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
