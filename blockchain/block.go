package blockchain

import transactions "github.com/Sucks-To-Suck/LuncheonNetwork/transactions"

type Block struct {
	blockHash   []byte
	prevHash    []byte
	maxWeight   uint32
	blockHeight uint32
	nonce       uint32

	BlockRewardTx transactions.PLuX
	basicTxs      []transactions.BLuX
	advancedTxs   []transactions.ALuX
}

func (b *Block) ConstructBlock() {

	b.BlockRewardTx.SetLuckyMiner([]byte("Joshua Decker"))
	b.BlockRewardTx.SetBlockReward(200)
	b.BlockRewardTx.GetWeight()
}
