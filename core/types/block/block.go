package block

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/core/events/txs"
	"github.com/Sucks-To-Suck/LuncheonNetwork/core/version"
	"github.com/Sucks-To-Suck/LuncheonNetwork/util/byteUtil"
)

// Represents the blocks on the blockchain.
// Each instance of this struct is its own block.
type Block struct {
	SoftwareVersion string
	PrevHash        string
	PackedTarget    uint32
	Miner           string

	MerkleRoot string
	Txs        []txs.LuTx

	Nonce      uint32 // Used to allow multiple workers to start at the same time, and have different hashes
	ExtraNonce uint32
	Timestamp  uint64

	BlockHash string
}

// Returns the newly created block.
func NewBlock(blockMinerId string) *Block {

	return new(Block)
}

// Add a mempool of transactions to the block.
// This mempool will be assumed to be the transactions chosen for the block, not just all transactions.
// Returns nothing.
func (b *Block) AddTxs(pool txs.Mempool) {

	// No Transactions
	if len(pool.Txs) == 0 {

		b.MerkleRoot = "NoTxs"

		return
	}

	b.Txs = append(b.Txs, pool.Txs...)

	b.MerkleRoot = b.GetMerkleRoot()
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

// Returns the bytes of the block needed for mining.
// Much faster and is also smaller than just getting the raw bytes of the block.
func (b *Block) MiningBytes() []byte {

	// Get the bytes of all information needed for mining.
	softwareVersion := version.GetVersionBytes()
	prevBlockHash, _ := hex.DecodeString(b.PrevHash)
	merkleRoot, _ := hex.DecodeString(b.MerkleRoot)
	timestamp := byteUtil.Uint64toB(b.Timestamp)
	packedTargetBytes := byteUtil.Uint32toB(b.PackedTarget)
	nonceBytes := byteUtil.Uint32toB(b.Nonce)
	extraNonce := byteUtil.Uint32toB(b.ExtraNonce)

	// Collapse them together into one byte array so it can be returned.
	softwareVersion = append(softwareVersion, prevBlockHash...)
	softwareVersion = append(softwareVersion, merkleRoot...)
	softwareVersion = append(softwareVersion, packedTargetBytes...)
	softwareVersion = append(softwareVersion, timestamp...)
	softwareVersion = append(softwareVersion, nonceBytes...)
	softwareVersion = append(softwareVersion, extraNonce...)

	return softwareVersion
}
