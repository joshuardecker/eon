package threader

import (
	"errors"
	"math/big"

	"github.com/Sucks-To-Suck/Eon/types/block"
)

var (
	ERR_INVALID       = errors.New("Invalid block submitted to threader")
	ERR_INVALID_PROOF = errors.New("Invalid proof type of the block")
)

// Takes the given block and attempts to submit the block. If the block is invalid, an error is returned.
func (t *Threader) SubmitBlock(b *block.Block, threadId *big.Int) error {

	switch b.Header().Proof() {

	// Case of the block being Proof of Authority.
	case "a":

		// If the block is invalid, return so.
		if !t.engine.VerifyBlock(b) {

			return ERR_INVALID
		}

		// Store the valid block in the given threads DB.
		bHash := b.Hash()
		err := t.threads[threadId.Uint64()].DB.Put(bHash.GetBytes(), b.Bytes(), nil)

		return err
	}

	return ERR_INVALID_PROOF
}
