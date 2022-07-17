package blockchain

import "github.com/Sucks-To-Suck/LuncheonNetwork/luncheon/transactions"

type Mempool struct {
	basicTxs    []transactions.BLuX
	advancedTxs []transactions.ALuX
}

func (m *Mempool) AddBasicToPool(tx transactions.BLuX) {

	m.basicTxs = append(m.basicTxs, tx)
}
