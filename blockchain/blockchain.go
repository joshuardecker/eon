package blockchain

import "errors"

// The blockchain struct that will be the chain of blocks.
type Blockchain struct {
	Blocks []Block

	height    uint32
	maxWeight uint32
}

func (b *Blockchain) InitBlockchain() error {

	if len(b.Blocks) != 0 || b.height != 0 {

		return errors.New("cannot init a blockchain when one already exists on this struct")
	}

	b.maxWeight = 1000000 // 1,000,000 aka one MegaByte, just a little bigger as some values are excluded from the weight factoring
	b.height = 0

	// Create the genisis block

	return nil
}
