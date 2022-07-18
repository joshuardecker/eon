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
	BlockHash    []byte
	PrevHash     []byte
	MaxWeight    uint32
	BlockHeight  uint32
	PackedTarget uint32

	Nonce uint32

	BlockRewardTx transactions.PLuX
	BasicTxs      []transactions.BLuX
	AdvancedTxs   []transactions.ALuX
}

// Constucts a block TODO: make this better when transactions made
func (b *Block) ConstructBlock() {

	b.BlockRewardTx.SetLuckyMiner(ellip.PubKeyHashStr())
	b.BlockRewardTx.SetBlockReward(200)
	b.BlockRewardTx.GetWeight()

	b.PackedTarget = 0x1dffffff
}

// Takes all of the block data and gets its bytes from memory. Returns those bytes.
func (b *Block) ParseBlockToBytes() []byte {

	// Gets the bytes that the block takes in memory
	blockAsBytes, jsonErr := json.Marshal(b)

	if jsonErr != nil {

		panic(jsonErr)
	}

	return blockAsBytes
}

// Sets the block hash, used by the miner. Returns any errors.
func (b *Block) SetBlockHash(inputHash []byte) error {

	// Is there no / blank input?
	if inputHash == nil {

		return errors.New("please input a non empty hash")
	}

	// Set the input hash
	b.BlockHash = inputHash

	return nil
}

// Prints the block Json
func (b *Block) PrintBlock() {

	// Gets the bytes that the block takes in memory
	blockJson, _ := json.Marshal(b)

	// Prints the json string of the block
	fmt.Println(string(blockJson))
}
