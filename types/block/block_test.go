package block

import (
	"math/big"
	"testing"

	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
	"github.com/Sucks-To-Suck/Eon/eocrypt/curve"
	"github.com/Sucks-To-Suck/Eon/helper"
)

func TestBlock(t *testing.T) {

	// Header variables:
	HEAD_HASH := eocrypt.HashBytes([]byte("My cool string!"))
	private, _ := curve.GenerateKeys()
	COINBASE := curve.CompressPub(&private.PublicKey)
	MERKLE := eocrypt.HashBytes([]byte("Merkle"))
	DIFF := big.NewInt(1)
	GAS := gas.Gas(1)
	GAS_MAX := gas.Gas(1)
	TIME := helper.LocalTime()
	NONCE := uint(1)
	SIG := []byte("Totally a real sig")

	// The Header created.
	h := NewHead(*HEAD_HASH, COINBASE, *MERKLE, DIFF, GAS, GAS_MAX, TIME, NONCE, SIG)

	// The block created.
	b := NewBlock(h, nil, nil, TIME)

	// Lets test some functions to see if they break:

	// Get the bytes of the block.
	if b.Bytes() == nil {

		panic("Cant be nil!")
	}

	// Print the block.
	b.Print()
}
