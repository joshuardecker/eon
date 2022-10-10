package block

import (
	"fmt"
	"math/big"
	"time"

	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
)

// A Header defines all of the significant features of a block, such as its parent block, coinbase, gas limit and more.
type Header struct {
	parentHash eocrypt.Hash `json:"ParentHash"`
	coinbase   []byte       `json:"Coinbase"`
	merkleRoot eocrypt.Hash `json:"Merkle"`
	difficulty *big.Int     `json:"Diff"`
	gas        gas.Gas      `json:"GasUsed"`
	gasLim     gas.Gas      `json:"MaxGas"`
	time       time.Time    `json:"Time"`
	nonce      uint         `json:"Nonce"`
	signature  []byte       `json:"Signature"`
	threadId   *big.Int     `json:"Thread"`
}

// Creates and gives a Head with the given inputs.
func NewHead(ParentHash eocrypt.Hash, Coinbase []byte, Merkle eocrypt.Hash, Diff *big.Int,
	Gas gas.Gas, GasLim gas.Gas, Time time.Time, Nonce uint) *Header {

	return &Header{
		parentHash: ParentHash,
		coinbase:   Coinbase,
		merkleRoot: Merkle,
		difficulty: Diff,
		gas:        Gas,
		gasLim:     GasLim,
		time:       Time,
		nonce:      Nonce,
	}
}

// Returns the hash of the Header.
func (h *Header) Hash() *eocrypt.Hash {

	return eocrypt.HashInterface(
		[]interface{}{

			h.parentHash.GetBytes(),
			h.coinbase,
			h.merkleRoot.GetBytes(),
			h.difficulty.Bytes(),
			h.gas.Uint(),
			h.gasLim.Uint(),
			h.nonce,
			h.signature,
		},
	)
}

func (h *Header) ParentHash() eocrypt.Hash { return h.parentHash }
func (h *Header) Coinbase() []byte         { return h.coinbase }
func (h *Header) MerkleRoot() eocrypt.Hash { return h.merkleRoot }
func (h *Header) Difficulty() *big.Int     { return h.difficulty }
func (h *Header) Gas() gas.Gas             { return h.gas }
func (h *Header) GasLimit() gas.Gas        { return h.gasLim }
func (h *Header) Time() time.Time          { return h.time }
func (h *Header) Nonce() uint              { return h.nonce }
func (h *Header) Signature() []byte        { return h.signature }

func (h *Header) Print() {

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
	`, h.parentHash, h.coinbase, h.merkleRoot, h.difficulty.Bytes(), h.gas, h.gasLim, h.time, h.nonce)
}
