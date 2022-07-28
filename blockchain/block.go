package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
	"github.com/Sucks-To-Suck/LuncheonNetwork/utilities"
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
	block.SoftwareVersion = utilities.SoftwareVersion
	block.PrevHash = b.Blocks[b.GetHeight()].BlockHash
	block.PackedTarget = b.CalculatePackedTarget(uint(len(b.Blocks)))
	block.Miner = blockMinerId

	return *block
}

// This function adds a slice of tx to the block.
// Input is the tx slice.
// Returns a bool, true if the tx were added, false if not.
func (b *Block) AddTx(tx transactions.LuTx) bool {

	// If the block weight + the new tx total weight is greater than the max weight
	if (tx.GetWeight() + b.GetWeight()) > MaxWeight {

		return false
	}

	b.Txs = append(b.Txs, tx)

	b.MerkleRoot = b.GetMerkleRoot()

	return true
}

// This function simply removes a tx from the block.
// Input is the tx index.
// Returns nothing.
func (b *Block) RemoveTx(txIndex uint) {

	// If the tx does not exist
	if txIndex >= uint(len(b.Txs)) {

		return
	}

	b.Txs = append(b.Txs[:txIndex], b.Txs[txIndex+1:]...)
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

// Converts the block into its bytes,
// Returns the byte slice of the block.
func (b *Block) AsBytes() []byte {

	// Get the byte slice
	bAsBytes, err := json.Marshal(b)

	if err != nil {

		panic(err)
	}

	return bAsBytes
}

// Prints the block as a string in JSON format.
// Returns nothing.
func (b *Block) PrintBlock() {

	// Gets the bytes that the block takes in memory
	blockJson, _ := json.Marshal(b)

	// Prints the json string of the block
	fmt.Println(string(blockJson))
}
