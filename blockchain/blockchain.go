package blockchain

import (
	"github.com/GoblinBear/beson/types"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/utilities"
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
	genisisB.SoftwareVersion = utilities.SoftwareVersion
	genisisB.PrevHash = "CoolGenisisBLock"
	genisisB.PackedTarget = 0x1d0fffff

	// Get the main public key ready
	mainKeys := new(ellip.MainKey)

	// Adds the genisis reward miner (you)
	genisisB.Miner = mainKeys.GetPubKeyStr()

	// Adds the genisis block to the blockchain
	b.AddBlock(genisisB)

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
// Input is the block thats being added.
func (b *Blockchain) AddBlock(block *Block) {

	b.Blocks = append(b.Blocks, *block)
}

// This function removes the last block from the blockchain.
// Returns nothing.
func (b *Blockchain) RemoveBlock() {

	b.Blocks = append(b.Blocks[:b.GetHeight()], b.Blocks[b.GetHeight()+1:]...)
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

// Calculates the packed target of a block.
// Expects what the block number will be, not what the current highest block is.
// So if this is used to see what the target of a new block will be, input what block height it will be.
// Returns the packed target of the block.
func (b *Blockchain) CalculatePackedTarget(blockNumber uint) uint32 {

	if blockNumber > uint(len(b.Blocks)) {

		return 0
	}

	// If block time is 1 minute, this will happen once a week
	if blockNumber%10080 == 0 {

		unPacker := new(utilities.TargetUnpacker)
		packer := new(utilities.TargetPacker)
		byteUtil := new(utilities.ByteUtil)
		time := b.Blocks[blockNumber-1].Timestamp - b.Blocks[blockNumber-10080].Timestamp

		newMultiplier := 10080 / time

		// Convert this to a uint256
		bigNewMultiplier := *types.NewUInt256("0", 1)
		bigNewMultiplier.Set(byteUtil.Uint64toB(newMultiplier))

		// Convert the current target to uint256
		target := unPacker.Unpack(b.Blocks[blockNumber-1].PackedTarget)

		// Apply the multiplier to the current target to get the new target
		newTarget := target.Multiply(&bigNewMultiplier)

		return packer.PackTargetUint256(*newTarget)
	}

	return b.Blocks[blockNumber-1].PackedTarget
}
