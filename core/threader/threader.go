package threader

import (
	"crypto/ecdsa"
	"math/big"

	poa "github.com/Sucks-To-Suck/Eon/PoA"
	pob "github.com/Sucks-To-Suck/Eon/PoB"
	pow "github.com/Sucks-To-Suck/Eon/PoW"
	"github.com/Sucks-To-Suck/Eon/types/thread"
)

const (
	MIN_THREADS = 1 // Minimum amount of threads in the Threader.
)

// The Threader is the mastermind of the blockchains that are threaded together. It covers
// consensus / verification of the various threads, and can expand or shrink the number of
// threads deployed.
type Threader struct {

	// The three types of consensus engines that can be used by the threader.
	// All three are saved that way threads of different types can be verified concurrently.
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
