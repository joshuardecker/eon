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
