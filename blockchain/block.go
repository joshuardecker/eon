package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/client"
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

// Represents the blocks on the blockchain.
// Each instance of this struct is its own block.
type Block struct {
	SoftwareVersion string
	PrevHash        string
	PackedTarget    uint32
	Miner           string

	MerkleRoot string
	Txs        []transactions.LuTx

	Nonce     uint32
	Timestamp uint64

	BlockHash string
}

// Creates a new block.
// Only input is the mining address that will be rewarded if the block is solved.
// Returns the newly created block.
func (b *Blockchain) CreateBlock(blockMinerId string) Block {

	// If the blockchain has not been initialized, return an empty block
	if len(b.Blocks) == 0 {

		return Block{}
	}

	block := new(Block)

	// Pack in the block information
	block.SoftwareVersion = client.SoftwareVersion
	block.PrevHash = b.Blocks[b.GetHeight()].BlockHash
	block.PackedTarget = b.CalculatePackedTarget(uint(len(b.Blocks)))
	block.Miner = blockMinerId

	return *block
}

// This function adds a slice of tx to the block.
// Input is the tx slice.
// Returns a bool, true if the tx were added, false if not.
func (b *Block) AddTx(tx []transactions.LuTx) bool {

	txWeight := uint(0)

	// Get the total weight of all the input txs
	for txIndex := 0; txIndex < len(tx); txIndex += 1 {

		txWeight += tx[txIndex].GetWeight()
	}

	// If the block weight + the new tx total weight is greater than the max weight
	if (txWeight + b.GetWeight()) > MaxWeight {

		return false
	}

	b.Txs = append(b.Txs, tx...)

	b.MerkleRoot = b.GetMerkleRoot()

	return true
}

// Calculates the weight of the block.
// Returns a uint32 of the block weight.
func (b *Block) GetWeight() uint {

	blockAsBytes, jsonErr := json.Marshal(b)

	if jsonErr != nil {

		panic(jsonErr)
	}

	return uint(len(blockAsBytes))
}

// Prints the block as a string in JSON format.
// Returns nothing.
func (b *Block) PrintBlock() {

	// Gets the bytes that the block takes in memory
	blockJson, _ := json.Marshal(b)

	// Prints the json string of the block
	fmt.Println(string(blockJson))
}
