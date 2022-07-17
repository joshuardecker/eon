package blockchain

import "github.com/Sucks-To-Suck/LuncheonNetwork/luncheon/transactions"

type Block struct {
	blockHash   []byte
	prevHash    []byte
	maxWeight   uint32
	blockHeight uint32
	nonce       uint32

	blockRewardTx transactions.PLuX
	basicTxs      transactions.BLuX
	advancedTxs   transactions.ALuX
}
