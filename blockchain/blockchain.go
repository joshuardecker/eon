package blockchain

import (
	"errors"

	"github.com/Sucks-To-Suck/LuncheonNetwork/client"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
)

// The blockchain struct that will be the chain of blocks.
type Blockchain struct {
	Blocks []Block

	height uint
}

// 1,000,000 aka one MegaByte, just a little bigger as some values are excluded from the weight factoring
var MaxWeight uint = 1000000

// Inits the blockchain struct, including defining constants.
// Creates the genisis block.
// Returns if any errors occured.
func InitBlockchain() Blockchain {

	// Create a blockchain instance
	b := new(Blockchain)

	b.height = 0

	// Create the genisis block:
	genisisB := new(Block)

	// Manually sets the variables of the genisis block
	genisisB.SoftwareVersion = client.SoftwareVersion
	genisisB.PrevHash = "CoolGenisisBLock"
	genisisB.PackedTarget = 0x1d0fffff

	// Get the main public key ready
	mainKeys := new(ellip.MainKey)

	// Adds the genisis reward miner (you)
	genisisB.Miner = mainKeys.GetPubKeyStr()

	// Adds the genisis block to the blockchain
	b.AddBlock(*genisisB)

	return *b
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
func (b *Blockchain) GetBlockReward(height uint32) uint64 {

	halvings := height / 525600

	// If no halvings have happened
	if halvings == 0 {

		// The default block reward (the * 1000000 is to convert LNCH to LUNCHEON)
		return 200 * 1000000
	}

	// If 1 or more halvings have happened
	// The << operator here acts as an easy way to do "to the power of" or **
	// Does not work in substitute for 2**0
	return (200 / (2 << (halvings - 1))) * 1000000
}

// Updates and returns the height of the blockchain.
// Returns a uint32 of the blockchain height.
func (b *Blockchain) GetHeight() uint {

	b.height = uint(len(b.Blocks) - 1)

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
func (b *Blockchain) VerifyBlock(block Block) bool {

	return true
}

// This function gets a block at a specified index.
// Returns the block and true if this was successful.
// If the index is invalid, it will return a empty block and false.
func (b *Blockchain) GetBlock(blockNum uint) (Block, bool) {

	if blockNum > b.GetHeight() {

		return Block{}, false
	}

	return b.Blocks[blockNum], true
}
