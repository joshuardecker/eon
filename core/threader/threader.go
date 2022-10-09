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
	engine Engine

	// The threads of the threader.
	threads *[]thread.Thread

	// The length of the threads.
	// This is saved as a big int, that way we dont need to rely on the len() function to get the length of the threads array.
	// Relying on the len() function limits the number to an int, making storing it as a Big.Int useless.
	threadsLen *big.Int

	// A lock that prevents multiple threads from being made at the same time, which would give them the same ID.
	threadsLock *sync.Mutex

	// The private key used by the threader.
	priv *ecdsa.PrivateKey

	// The configuration that will be used by the threader.
	config *config.Config
}

// Creates a new threader with the given configuration, proofEngine, and the nodes private key.
func NewThreader(config config.Config, proofEngine Engine, private ecdsa.PrivateKey) *Threader {

	// Create the threader.
	t := new(Threader)

	// Lock in the proof engine.
	t.engine = proofEngine

	// Add a single Thread to start.
	t.AddThread(thread.NewThread(big.NewInt(0)))

	t.priv = &private
	t.config = &config

	return t
}

// Adds the given thread to the threader. No need to add an ID to the thread before calling this function.
func (t *Threader) AddThread(th *thread.Thread) {

	// Lock the threads lock while adding a new thread.
	t.threadsLock.Lock()

	// Unlock it when done.
	defer t.threadsLock.Unlock()

	// The id equals the length of the threads array, so if two threads exist, the new thread will have the id of "3".
	th.Id = t.threadsLen

	// Add 1 to the length of the threads.
	t.threadsLen.Add(t.threadsLen, big.NewInt(1))

	// Add the new thread to the array.
	*t.threads = append(*t.threads, *th)
}

// Loads the given thread based on id. If the thread does not exist, it will create a thread with that id.
func (t *Threader) LoadThread(id *big.Int) {

	// Load / create a thread with the given id.
	th := thread.NewThread(id)

	// Add the thread to the threader.
	t.AddThread(th)
}

// Loads all of the threads with the given largest id. Example: if 4 is provided as the id, this func
// will load thread with the id of '4', '3', '2', and '1'.
func (t *Threader) LoadAllThreads(id *big.Int) {

	// Loop through, starting at the highest thread id and looping loading / creation of the
	// threads until it hits loops until thread 0 (by default the only thread).
	for id.Cmp(big.NewInt(0)) >= 0 {

		// Load / create the thread.
		t.LoadThread(id)

		// Subtract 1 from the id so it can loop to 0.
		id.Sub(id, big.NewInt(1))
	}
}
