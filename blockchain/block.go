package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

// Represents the blocks on the blockchain.
// Each instance of this struct is its own block.
type Block struct {
	SoftwareVersion string
	PrevHash        string
	PackedTarget    uint32

	MerkleRoot string
	Txs        []transactions.LuTx

	Miner     string
	Nonce     uint32
	Timestamp uint64

	BlockHash string
}

// Creates a new block.
// The inputs are:
// This function returns:
func (b *Block) CreateBlock() {}

// Calculates the weight of the block.
// Returns a uint32 of the block weight.
func (b *Block) GetWeight() uint32 { return 0 }

// Prints the block as a string in JSON format.
// Returns nothing.
func (b *Block) PrintBlock() {

	// Gets the bytes that the block takes in memory
	blockJson, _ := json.Marshal(b)

	// Prints the json string of the block
	fmt.Println(string(blockJson))
}
