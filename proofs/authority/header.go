package poa

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
)

var (
	ERR_SIGN = errors.New("Could not sign the block")
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
	proof      string       `json:"Proof"` // Should be "a" "b" or "w" to signify an authority proof, burn proof, or work proof.
}

// Creates and gives a Head with the given inputs.
func NewHeader(ParentHash eocrypt.Hash, Coinbase []byte, Merkle eocrypt.Hash,
	Gas gas.Gas, GasLim gas.Gas, Time time.Time, ThreadId *big.Int) *PoAHeader {

	return &PoAHeader{
		parentHash: ParentHash,
		coinbase:   Coinbase,
		merkleRoot: Merkle,
		gas:        Gas,
		gasLim:     GasLim,
		time:       Time,
		threadId:   ThreadId,
		proof:      "a",
	}
}

// Returns the hash of the Header.
func (h *PoAHeader) Hash() eocrypt.Hash {

	return *eocrypt.HashInterface(
		[]interface{}{

			h.parentHash.Bytes(),
			h.coinbase,
			h.merkleRoot.Bytes(),
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
func (h *PoAHeader) Proof() string            { return h.proof }

// Sets the signature of the block to the given bytes.
// Does NOT actually sign the header, just saves the signature.
func (h *PoAHeader) Sign(sig []byte) {

	h.signature = sig
}

func (h *PoAHeader) Print() {

	fmt.Printf(`
	[
		Type: Proof of Authority
		Parent Hash: %x
		Coinbase: %x
		Merkle: %x
		Gas: %d
		Max Gas: %d	
		Time: %v
	]
	`, h.parentHash, h.coinbase, h.merkleRoot, h.gas, h.gasLim, h.time)
}

// Prints the header with the given block hash, aka printing the block.
func (h *PoAHeader) PrintBlock(bHash eocrypt.Hash) {

	fmt.Printf(`
	[
		Type: Proof of Authority
		Block Hash: %x

		Parent Hash: %x
		Coinbase: %x
		Merkle: %x
		Gas: %d
		Max Gas: %d	
		Time: %v
	]
	`, bHash, h.parentHash, h.coinbase, h.merkleRoot, h.gas, h.gasLim, h.time)
}
