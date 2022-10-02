package threader

import (
	"crypto/ecdsa"
	"math/big"
	"sync"

	"github.com/Sucks-To-Suck/Eon/core/config"
	"github.com/Sucks-To-Suck/Eon/types/thread"
)

// The Threader is the mastermind of the blockchains that are threaded together. It covers
// consensus / verification of the various threads, and can expand or shrink the number of
// threads deployed.
type Threader struct {

	// The Proof Engine of the threader can be multiple types, so here we declare it as a interface for flexability.
	ProofEngine interface{}

	// The threads of the threader.
	Threads *[]thread.Thread

	// The length of the threads.
	// This is saved as a big int, that way we dont need to rely on the len() function to get the length of the threads array.
	// Relying on the len() function limits the number to an int, making storing it as a Big.Int useless.
	ThreadsLen *big.Int

	// A lock that prevents multiple threads from being made at the same time, which would give them the same ID.
	ThreadsLock *sync.Mutex

	// The private key used by the threader.
	Priv *ecdsa.PrivateKey

	// The configuration that will be used by the threader.
	config *config.Config
}

// Creates a new threader with the given configuration, proofEngine, and the nodes private key.
func NewThreader(config config.Config, proofEngine interface{}, private ecdsa.PrivateKey) *Threader {

	// Create the threader.
	t := new(Threader)

	// Lock in the proof engine.
	t.ProofEngine = proofEngine

	// Add a single Thread to start.
	t.AddThread(thread.NewThread(big.NewInt(0)))

	t.Priv = &private
	t.config = &config

	return t
}

// Adds the given thread to the threader. No need to add an ID to the thread before calling this function.
func (t *Threader) AddThread(th *thread.Thread) {

	// Lock the threads lock while adding a new thread.
	t.ThreadsLock.Lock()

	// Unlock it when done.
	defer t.ThreadsLock.Unlock()

	// The id equals the length of the threads array, so if two threads exist, the new thread will have the id of "3".
	th.Id = t.ThreadsLen

	// Add 1 to the length of the threads.
	t.ThreadsLen.Add(t.ThreadsLen, big.NewInt(1))

	// Add the new thread to the array.
	*t.Threads = append(*t.Threads, *th)
}
