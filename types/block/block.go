package block

import (
	"math/big"
	"time"

	"github.com/Sucks-To-Suck/LuncheonNetwork/core/gas"
	"github.com/Sucks-To-Suck/LuncheonNetwork/helper"
)

// Headers define many features of the block, such as its parent or the set difficulty of the block.
// Used as the standard in Eon.
type Head struct {
	ParentHash helper.Hash `json:"ParentHash"`
	Coinbase   []byte      `json:"Coinbase"` // TODO: update when ellip is updated.
	MerkleRoot helper.Hash `json:"Merkle"`
	Difficulty *big.Int    `json:"Diff"`
	Gas        gas.Gas     `json:"GasUsed"`
	MaxGas     gas.Gas     `json:"MaxGas"`
	Time       time.Time   `json:"Time"`
	Nonce      uint64      `json:"Nonce"`
	ExtraNonce uint64      `json:"ExtraNonce"`
}

// Light Headers are used when operating Eon on a private network and the features of the normal Header are unnessesary.
// Example, on a PoA (proof of authority) based network, things like difficulty are simply not needed.
type LightHead struct {
	ParentHash helper.Hash `json:"ParentHash"`
	Coinbase   []byte      `json:"Coinbase"` // TODO: update when ellip is updated.
	MerkleRoot helper.Hash `json:"Merkle"`
	Time       time.Time   `json:"Time"`
}

// Blocks are a storage of data transactions. They have an identifying hash, as well as Time created and received.
// Blocks also have a normal Head by default.
type Block struct {
	Head         *Head    `json:"Head"`
	Uncles       *[]Head  `json:"Uncles"`
	Transactions []string `json:"Txs"` // TODO: update when tx is made.

	BlockHash helper.Hash `json:"Hash"`
	Received  time.Time   `json:"Received"`
}

// Light blocks are used when wanting to use Light Heads on those blocks, which is mainly in a PoA (proof of authority) setting.
// Remains the same compared to normal blocks, in that data transactions are stored, along with a blockhash and time of creation and reception.
type LightBlock struct {
	Head         *LightHead   `json:"Head"`
	Uncles       *[]LightHead `json:"Uncles"`
	Transactions []string     `json:"Txs"` // TODO: update when tx is made.

	BlockHash helper.Hash `json:"Hash"`
	Received  time.Time   `json:"Received"`
}

// ****
// Head:

// Creates and gives a Head with the given inputs.
func NewHead(ParentHash helper.Hash, Coinbase []byte, Merkle helper.Hash, Diff *big.Int, // TODO: Update when ellip updated
	Gas gas.Gas, MaxGas gas.Gas, Time time.Time, Nonce uint64, ExtraNonce uint64) *Head {

	h := new(Head)

	h.SetParentHash(ParentHash)
	h.SetCoinbase(Coinbase)
	h.SetMerkle(Merkle)
	h.SetDiff(*Diff)
	h.SetGas(Gas)
	h.SetMaxGas(MaxGas)
	h.SetTime(Time)
	h.SetNonce(Nonce)
	h.SetExtraNonce(ExtraNonce)

	return h
}

// Returns the hash of the Header.
func (h *Head) Hash() *helper.Hash {

	return helper.HashInterface(
		[]interface{}{

			h.ParentHash,
			h.Coinbase,
			h.MerkleRoot,
			h.Difficulty,
			h.Gas,
			h.MaxGas,
			h.Nonce,
		},
	)
}

func (h *Head) SetParentHash(hash helper.Hash) {

	h.ParentHash = hash
}

func (h *Head) SetCoinbase(c []byte) { // TODO: Update when ellip updated

	h.Coinbase = c
}

func (h *Head) SetMerkle(hash helper.Hash) {

	h.MerkleRoot = hash
}

// Note about this function: none of these funcs use pointer inputs, that way data in them ant be changed unless specifically
// set by a function for safety reasons, but making a copy of the Big Int may be costly in very high throughput systems.
func (h *Head) SetDiff(b big.Int) {

	h.Difficulty = &b
}

func (h *Head) SetGas(g gas.Gas) {

	h.Gas = g
}

func (h *Head) SetMaxGas(g gas.Gas) {

	h.Gas = g
}

func (h *Head) SetTime(t time.Time) {

	h.Time = t
}

func (h *Head) SetNonce(n uint64) {

	h.Nonce = n
}

func (h *Head) SetExtraNonce(en uint64) {

	h.ExtraNonce = en
}

// Head:
// ****

// ****
// Light Head:

// Creates and gives a Light Header with the given inputs.
func NewLightHead(ParentHash helper.Hash, Coinbase []byte, Merkle helper.Hash, Time time.Time) *LightHead {

	lh := new(LightHead)

	lh.SetParentHash(ParentHash)
	lh.SetCoinbase(Coinbase)
	lh.SetMerkle(Merkle)
	lh.SetTime(Time)

	return lh
}

// Returns the hash of the Light Header.
func (h *LightHead) Hash() *helper.Hash {

	return helper.HashInterface(
		[]interface{}{

			h.ParentHash,
			h.Coinbase,
			h.MerkleRoot,
		},
	)
}

func (h *LightHead) SetParentHash(hash helper.Hash) {

	h.ParentHash = hash
}

func (h *LightHead) SetCoinbase(c []byte) { // TODO: Update when ellip updated

	h.Coinbase = c
}

func (h *LightHead) SetMerkle(hash helper.Hash) {

	h.MerkleRoot = hash
}

func (h *LightHead) SetTime(t time.Time) {

	h.Time = t
}

// Light Head:
// ****

// ****
// Block:

// Block:
// ****

// ****
// Light Block:

// Light Block:
// ****
