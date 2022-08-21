package worker

import (
	"bytes"
	"encoding/hex"
	"math/big"

	"github.com/Sucks-To-Suck/LuncheonNetwork/core/events/block"
	"github.com/Sucks-To-Suck/LuncheonNetwork/util"
	"golang.org/x/crypto/sha3"
)

// The worker is a single instance that will try to solve a block with the given parameters.
type Worker struct {
	currentHash []byte
	target      []byte

	Mining    bool
	HashCount uint32

	utilTime util.Time
}

// Creates and returns a new worker.
func New(target *big.Int) *Worker {

	w := new(Worker)

	// Fill data
	w.target = target.Bytes()
	w.Mining = true

	return w
}

// Tells the worker to start the mining process.
// Input the block being worked on, an extra nonce, and this will search until a solution is found.
// If one is found, it is returned as a hex string.
func (w *Worker) Start(b *block.Block, extraNonce uint32) string {

	// Set the extra nonce, so workers are easily making unique hashes.
	b.ExtraNonce = extraNonce

	// The actual mining process. Only continues if Mining is set to true, so if its false, the worker will stop.
	for b.Nonce = 0; w.Mining; b.Nonce++ {

		// Set the timestamp in the block.
		b.Timestamp = w.utilTime.CurrentUnix()

		// Create the hash size (32 bytes).
		w.currentHash = make([]byte, 32)

		// Hash the data.
		sha3.ShakeSum256(w.currentHash, b.MiningBytes())

		// Was the solution found?
		if bytes.Compare(w.currentHash, w.target) != 1 {

			return hex.EncodeToString(w.currentHash)
		}
	}

	return ""
}
