package blockchain

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

// Blocks for the chain, each one containing its own transaction data to support the network.
type Block struct {
	// Things known before the new block
	SoftwareVersion uint32
	PrevHash        string
	PackedTarget    uint32

	// Things calculated after creation
	BlockHash string
	Nonce     uint32
	Timestamp uint64

	Weight uint32 // Represents the Weight of all transactions, and excludes everything else (for obvious reason of it would changed the hash)

	// Tx's (and their merkle root)
	PLuX       transactions.PLuX
	BLuX       []transactions.BLuX
	ALuX       []transactions.ALuX
	MerkleRoot string
}

// Inits a block
func (b *Block) InitBlock(prevHash string, packedTarget uint32) {

	// Init the new block
	b.PrevHash = prevHash
	b.PackedTarget = packedTarget
}

// Calculates the weight of the block
func (b *Block) GetWeight() {

	// Gets the total weight of the reward transaction
	byteArray, jsonErr := json.Marshal(b.PLuX)

	if jsonErr != nil {
		panic(jsonErr)
	}

	b.Weight += uint32(len(byteArray))

	// Gets the total weight of the basic transactions
	byteArray, jsonErr = json.Marshal(b.BLuX)

	if jsonErr != nil {
		panic(jsonErr)
	}

	b.Weight += uint32(len(byteArray))

	// Gets the total weight of the advanced transactions
	byteArray, jsonErr = json.Marshal(b.ALuX)

	if jsonErr != nil {
		panic(jsonErr)
	}

	b.Weight += uint32(len(byteArray))
}

// Sets the block hash found by the miner. Returns any errors.
func (b *Block) SetBlockHash(inputHash []byte) error {

	// Is there no / blank input?
	if inputHash == nil {

		return errors.New("please input a non empty hash")
	}

	// Set the input hash
	b.BlockHash = hex.EncodeToString(inputHash)

	return nil
}

// Nonce is already handled by the miner //

// Takes in a unix timestamp. Blocks will use unix in seconds.
func (b *Block) SetTimestamp(unixTime uint64) {

	b.Timestamp = unixTime
}

// Adds the PLuX tranaction to the block. Returns errors.
func (b *Block) AddPLuX(p transactions.PLuX) error {

	b.PLuX = p

	return nil
}

// Simply adds the basic transaction to the list. TODO: add BLuX check
func (b *Block) AddBLuX(bl transactions.BLuX) error {

	b.BLuX = append(b.BLuX, bl)

	return nil
}

func (b *Block) AddALuX(a transactions.ALuX) error {

	b.ALuX = append(b.ALuX, a)

	return nil
}

//**********
//
//**********

// Returns the byte array of the block struct.
// Will be called with all data spots filled except for BlockHash, which will be left empty until one is found.
func (b *Block) ParseBlockToBytes() []byte {

	// Gets the bytes that the block takes in memory
	blockAsBytes, jsonErr := json.Marshal(b)

	// If an error happened
	if jsonErr != nil {

		panic(jsonErr)
	}

	return blockAsBytes
}

// Prints the block Json string.
func (b *Block) PrintBlock() {

	// Gets the bytes that the block takes in memory
	blockJson, _ := json.Marshal(b)

	// Prints the json string of the block
	fmt.Println(string(blockJson))
}
