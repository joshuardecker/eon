package poa

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/Sucks-To-Suck/Eon/core/config"
	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
	"github.com/Sucks-To-Suck/Eon/eocrypt/curve"
	"github.com/Sucks-To-Suck/Eon/types/block"
)

func TestConsensus(t *testing.T) {

	// Tests the consensus of Proof of Authority Blocks.
	// Goes from block creation, to consensus.

	// Generate the parameters of the Authority engine:
	config := new(config.Config)
	key, _ := curve.GenerateKeys()

	config.TrustedKey = &key.PublicKey

	ae := NewAuthorityEngine(*config, *key)

	// Creates the Proof of Authority Header:
	parentHash := eocrypt.HashBytes([]byte("From this block"))
	coinbase := curve.CompressPub(&key.PublicKey)
	merkleRoot := eocrypt.HashBytes([]byte("Kai Morts transactions"))
	gasUsed := gas.Gas(1)
	gasLim := gas.Gas(1)
	hTime := time.Now()
	id := big.NewInt(1)

	h := NewHeader(*parentHash, coinbase, *merkleRoot, gasUsed, gasLim, hTime, id)

	// Creates the Proof of Authority Block:

	b := block.NewBlock(h, nil, nil, hTime)

	b.CalcHash() // Calculate and store the hash of the block.

	// Lets test printing:
	b.Print()

	// Lets test consensus:
	ae.ValidateBlock(b)

	fmt.Println(ae.VerifyBlock(b)) // Should Print 'true'.
}
