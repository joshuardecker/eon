package poa

import (
	"fmt"
	"math/big"
	"time"

	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
)

// The Header that Proof of Authority blocks use.
// Only contains perameters that PoA requires.
type PoAHeader struct {
	parentHash eocrypt.Hash `json:"ParentHash"`
	coinbase   []byte       `json:"Coinbase"`
	merkleRoot eocrypt.Hash `json:"Merkle"`
	gas        gas.Gas      `json:"GasUsed"`
	gasLim     gas.Gas      `json:"MaxGas"`
	time       time.Time    `json:"Time"`
	signature  []byte       `json:"Signature"`
	threadId   *big.Int     `json:"Thread"`
}

// Creates and gives a Head with the given inputs.
func NewHead(ParentHash eocrypt.Hash, Coinbase []byte, Merkle eocrypt.Hash, Diff *big.Int,
	Gas gas.Gas, GasLim gas.Gas, Time time.Time, Nonce uint) *PoAHeader {

	return &PoAHeader{
		parentHash: ParentHash,
		coinbase:   Coinbase,
		merkleRoot: Merkle,
		gas:        Gas,
		gasLim:     GasLim,
		time:       Time,
	}
}

// Returns the hash of the Header.
func (h *PoAHeader) Hash() *eocrypt.Hash {

	return eocrypt.HashInterface(
		[]interface{}{

			h.parentHash.GetBytes(),
			h.coinbase,
			h.merkleRoot.GetBytes(),
			h.gas.Uint(),
			h.gasLim.Uint(),
			h.signature,
		},
	)
}

func (h *PoAHeader) ParentHash() eocrypt.Hash { return h.parentHash }
func (h *PoAHeader) Coinbase() []byte         { return h.coinbase }
func (h *PoAHeader) MerkleRoot() eocrypt.Hash { return h.merkleRoot }
func (h *PoAHeader) Gas() gas.Gas             { return h.gas }
func (h *PoAHeader) GasLimit() gas.Gas        { return h.gasLim }
func (h *PoAHeader) Time() time.Time          { return h.time }
func (h *PoAHeader) Signature() []byte        { return h.signature }

func (h *PoAHeader) Print() {

	fmt.Printf(`
	[
		Parent Hash: %x
		Coinbase: %x
		Merkle: %x
		Difficulty: %d
		Gas: %d
		Max Gas: %d	
		Time: %v
		Nonce: %v
	]
	`, h.parentHash, h.coinbase, h.merkleRoot, h.gas, h.gasLim, h.time)
}
