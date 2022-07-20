package blockchain

import (
	"errors"

	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

// The blockchain struct that will be the chain of blocks.
type Blockchain struct {
	Blocks []Block

	height    uint32
	maxWeight uint32
}

// Inits the blockchain struct, including defining constants.
// Returns if any errors occured.
func (b *Blockchain) InitBlockchain() error {

	if len(b.Blocks) != 0 || b.height != 0 {

		return errors.New("cannot init a blockchain when one already exists on this struct")
	}

	b.maxWeight = 1000000 // 1,000,000 aka one MegaByte, just a little bigger as some values are excluded from the weight factoring
	b.height = 0

	// Create the genisis block

	return nil
}

// This function creates the blockreward transaction.
// It will use your main key as the receiver.
// Input is the current block reward.
// Returns the created transaction.
func (b *Blockchain) CreateBlockRewardTx() transactions.LuTx {

	tx := new(transactions.LuTx)
	key := new(ellip.MainKey)

	pubKey := key.MainKeyHash()
	bReward := string(b.GetBlockReward())

	sc := "PUBKH " + pubKey + " "
	sc = "AMT " + bReward + " "
	sc = "TIML 25"

	tx.AddScriptStr(sc, false)

	return *tx
}

// Returns the current block reward.
// Just for some context, the average blocktime shoots for 1 minute.
// The blockchain reward will target to half once per year in Luncheon 1.0.
// Block reward starts at 200 per block.
// This means every 525600 blocks, the reward halves.
// The current code also makes it so the blockchain rewards besides tx fees
// Will fully dry-up in 7 years, the first block of year 8 will have zero reward.
// The total amount of coins that can exist is 208,663,200, which means 10 of these coins
// can be considered as rare, in terms of total in existance, as 1 btc.
func (b *Blockchain) GetBlockReward() uint64 {

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

// Updates and returns the height of the blockchain.
// Returns a uint32 of the blockchain height.
func (b *Blockchain) GetHeight() uint32 {

	b.height = uint32(len(b.Blocks) - 1)

	return b.height
}

// This function adds a block to the blockchain.
// Inputs are the block being added.
// Returns any errors.
func (b *Blockchain) AddBlock(block Block) error {

	if !b.VerifyBlock(block) {

		return errors.New("block is invalid")
	}

	b.Blocks = append(b.Blocks, block)

	return nil
}

// This function removes the last block from the blockchain.
// Returns nothing.
func (b *Blockchain) RemoveBlock() {

	b.Blocks = append(b.Blocks[:b.GetHeight()], b.Blocks[b.GetHeight()+1:]...)
}

// Determines whether the block is valid.
// Returns a bool, true if valid, and false if invalid.
func (b *Blockchain) VerifyBlock(block Block) bool { return true }
