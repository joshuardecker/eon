package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

// Represents the blocks on the blockchain.
// Each instance of this struct is its own block.
type Block struct {
	SoftwareVersion uint32
	PrevHash        string
	PackedTarget    uint32

	Nonce     uint32
	Timestamp uint64

	BlockHash string

	MerkleRoot string
	Txs        []transactions.LuTx
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

// Determines whether the block is valid.
// Returns a bool, true if valid, and false if invalid.
func (b *Blockchain) verifyBlock(block Block) bool { return true }
