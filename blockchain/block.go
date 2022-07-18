package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

// Blocks for the chain, each one containing its own transaction data to support the network.
type Block struct {
	blockHash    []byte
	prevHash     []byte
	maxWeight    uint32
	blockHeight  uint32
	packedTarget uint32

	Nonce uint32

	BlockRewardTx transactions.PLuX
	basicTxs      []transactions.BLuX
	advancedTxs   []transactions.ALuX
}

func (b *Block) ConstructBlock() {

	b.BlockRewardTx.SetLuckyMiner(ellip.PubKeyHashStr())
	b.BlockRewardTx.SetBlockReward(200)
	b.BlockRewardTx.GetWeight()

	b.packedTarget = 0x1dffffff
}

// Takes all of the block data and gets its bytes from memory. Returns those bytes.
func (b *Block) ParseBlockToBytes() []byte {

	blockAsBytes, jsonErr := json.Marshal(*b)

	if jsonErr != nil {
		panic(jsonErr)
	}

	return blockAsBytes
}

func (b *Block) SetBlockHash(inputHash []byte) error {

	if inputHash == nil {

		return errors.New("please input a non empty hash")
	}

	b.blockHash = inputHash

	return nil
}

func (b *Block) PrintBlock() {

	blockJson, _ := json.Marshal(b)

	fmt.Println(string(blockJson))
}
