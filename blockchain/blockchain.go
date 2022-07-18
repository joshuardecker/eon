package blockchain

import (
	"errors"
)

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

// Returns the current block reward.
// Just for some context, the average blocktime shoots for 1 minute.
// The blockchain reward will target to half once per year in Luncheon 1.0.
// Block reward starts at 200 per block.
// This means every 525,600 blocks, the reward halves.
// The current code also makes it so the blockchain rewards besides tx fees
// Will fully dry-up in 7 years, the first block of year 8 will have zero reward.
func (b *Blockchain) GetBlockReward() uint64 {

	b.height = 4204800
	halvings := b.height / 525600

	// If no halvings have happened
	if halvings == 0 {

		// The default block reward
		return 200
	}

	// If 1 or more halvings have happened
	// The << operator here acts as an easy way to do "to the power of" or **
	// Does not work in substitute for 2**0
	return 200 / (2 << (halvings - 1))
}
