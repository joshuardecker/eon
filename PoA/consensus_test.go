package poa

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/Sucks-To-Suck/Eon/core/config"
	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
	"github.com/Sucks-To-Suck/Eon/eocrypt/curve"
	"github.com/Sucks-To-Suck/Eon/helper"
	"github.com/Sucks-To-Suck/Eon/types/block"
)

func TestConsensus(t *testing.T) {

	// Generate the parameters of the Authority engine.
	config := new(config.Config)
	key, _ := curve.GenerateKeys()

	// Give the config a trusted key (in this case the generated one).
	config.TrustedKey = &key.PublicKey

	// *****
	// Creation of the block:

	// Header variables:
	HEAD_HASH := eocrypt.HashBytes([]byte("My cool string!"))
	COINBASE := curve.CompressPub(&key.PublicKey)
	MERKLE := eocrypt.HashBytes([]byte("Merkle"))
	DIFF := big.NewInt(1)
	GAS := gas.Gas(1)
	GAS_MAX := gas.Gas(1)
	TIME := helper.LocalTime()
	NONCE := uint(1)

	// The Header created.
	h := block.NewHead(*HEAD_HASH, COINBASE, *MERKLE, DIFF, GAS, GAS_MAX, TIME, NONCE)

	// The block created.
	b := block.NewBlock(h, nil, nil, TIME)

	// Creation of the block:
	// *****

	// Create the engine.
	ae := NewAuthorityEngine(*config, *key)

	// Validate the block.
	ae.ValidateBlock(b)

	// Check whether the block is valid (should print 'true').
	fmt.Println(ae.VerifyBlock(b))
}
