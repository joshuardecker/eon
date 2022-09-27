package threader

import (
	"crypto/ecdsa"
	"math/big"

	poa "github.com/Sucks-To-Suck/Eon/PoA"
	pob "github.com/Sucks-To-Suck/Eon/PoB"
	pow "github.com/Sucks-To-Suck/Eon/PoW"
	"github.com/Sucks-To-Suck/Eon/types/thread"
)

// The Threader is the mastermind of the blockchains that are threaded together. It covers
// consensus / verification of the various threads, and can expand or shrink the number of
// threads deployed.
type Threader struct {

	// Saved engines to shift between.
	PoAEngine *poa.AuthorityEngine
	PoBEngine *pob.BurnEngine
	PoWEngine *pow.WorkEngine

	// The threads of the Threader.
	Threads *[]thread.Thread

	// The chain id of this Threader.
	Id *big.Int

	// The public key used by the Threader.
	Pub *ecdsa.PublicKey
}
