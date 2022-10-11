package block

import (
	"time"

	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
)

// A Header is a flexable type that represents the blocks parameters and essential info.
// The variation in Headers is because of the variation in Proofs (Authority, Burn, and Work).
// All of those Proof types have different Header requirements, so a common interface is made to represent them all.
type Header interface {

	// Functions that retrieve information from the Headers:
	Hash() eocrypt.Hash

	ParentHash() eocrypt.Hash

	Coinbase() []byte

	MerkleRoot() eocrypt.Hash

	Gas() gas.Gas

	GasLimit() gas.Gas

	Time() time.Time

	Signature() []byte

	Proof() string

	// Functions that modify the in block data:
	Sign(sig []byte)

	// Misc functions:
	Print()

	PrintBlock(bHash eocrypt.Hash)
}
