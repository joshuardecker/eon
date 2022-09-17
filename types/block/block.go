package block

import (
	"math/big"
	"time"

	"github.com/Sucks-To-Suck/LuncheonNetwork/core/gas"
	"github.com/Sucks-To-Suck/LuncheonNetwork/helper"
)

// Headers define many features of the block. Lighter blocks may need a header.
// This is depending on how this application is used,
type Head struct {
	ParentHash helper.Hash `json:"ParentHash"`
	Coinbase   []byte      `json:"Coinbase"` // TODO: update when ellip is updated.
	MerkleRoot helper.Hash `json:"Merkle"`
	Difficulty *big.Int    `json:"Diff"`
	Gas        gas.Gas     `json:"GasUsed"`
	MaxGas     gas.Gas     `json:"MaxGas"`
	Nonce      uint64      `json:"Nonce"`
}

// Light Headers are used when operating Eon on a private network and the features of the normal Header are unnessesary.
type LightHead struct {
	ParentHash helper.Hash `json:"ParentHash"`
	Coinbase   []byte      `json:"Coinbase"` // TODO: update when ellip is updated.
}

// Blocks are a storage of data transactions. They have an identifying hash, as well as Time created and received.
// Heads and uncle headers may be unnesesary depending on the applications use of Eon.
type Block struct {
	Head         *Head    `json:"Head"`
	Uncles       *[]Head  `json:"Uncles"`
	Transactions []string `json:"Txs"` // TODO: update when tx is made.

	BlockHash helper.Hash `json:"Hash"`
	Time      time.Time   `json:"Time"`
	Received  time.Time   `json:"Received"`
}

// ****
// Head:

// Creates and gives a Head with the given inputs.
func NewHead(ParentHash helper.Hash, Coinbase []byte, Merkle helper.Hash, Diff *big.Int, // TODO: Update when ellip updated
	Gas gas.Gas, MaxGas gas.Gas, Nonce uint64) *Head {

	h := new(Head)

	h.ParentHash = ParentHash
	h.Coinbase = Coinbase
	h.MerkleRoot = Merkle
	h.Difficulty = Diff
	h.Gas = Gas
	h.MaxGas = MaxGas
	h.Nonce = Nonce

	return h
}

// Creates and gives a Light Header with the given inputs.
func NewLightHead(ParentHash helper.Hash, Coinbase []byte) *LightHead {

	lh := new(LightHead)

	lh.ParentHash = ParentHash
	lh.Coinbase = Coinbase

	return lh
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

func (h *Head) SetDiff(b *big.Int) {

	h.Difficulty = b
}

func (h *Head) SetGas(g gas.Gas) {

	h.Gas = g
}

func (h *Head) SetMaxGas(g gas.Gas) {

	h.Gas = g
}

func (h *Head) SetNonce(n uint64) {

	h.Nonce = n
}

// Head:
// ****
