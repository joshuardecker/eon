package blockchain

import (
	"encoding/binary"
	"encoding/json"
	"os"

	"github.com/GoblinBear/beson/types"
	"github.com/Sucks-To-Suck/LuncheonNetwork/core/events/block"
	"github.com/Sucks-To-Suck/LuncheonNetwork/crypto/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/util"
)

// The blockchain struct that will be the chain of blocks.
type Blockchain struct {
	Blocks []block.Block
}

// 1,000,000 aka one MegaByte, just a little bigger as some values are excluded from the weight factoring
var MaxWeight uint = 1000000

// Inits the blockchain struct, including defining constants.
// Creates the genisis block.
// Returns if any errors occured.
func InitBlockchain() *Blockchain {

	// Create a blockchain instance
	b := new(Blockchain)

	// Create the genisis block:
	genisisB := new(block.Block)

	// Manually sets the variables of the genisis block
	genisisB.SoftwareVersion = util.SoftwareVersion
	genisisB.PrevHash = "CoolGenisisBlock"
	genisisB.PackedTarget = 0x1e050000

	// Get the main public key ready
	mainKeys := new(ellip.MainKey)

	// Adds the genisis reward miner (you)
	genisisB.Miner = mainKeys.GetPubKeyStr()

	// Adds the genisis block to the blockchain
	b.AddBlock(genisisB)

	return b
}

// Returns the current block reward.
// Just for some context, the average blocktime shoots for 5 seconds.
// The blockchain reward will reduce by 20% every 2 years in Luncheon 1.0.
// Block reward starts at 5 per block.
// This means every 6,307,200 blocks, the reward reduces by 20%.
func (b *Blockchain) GetBlockReward(height uint32) uint64 {

	// Amount of times the block reward is being reduced
	reduces := height / 6307200

	reward := uint64(5 * 1000000) // Default reward

	// Loops and applys block reward reduction based on the amount of reductions that have been passed.
	// Based on block height, as calculated above.
	for loop := uint32(0); loop < reduces; loop += 1 {

		reward *= (8 * 100000) // It is * 100K because 8 * 100K = 0.8 * 1M, and the *0.8 is the 20% block reduction
		reward /= 1000000
	}

	return reward
}

// Returns a uint of the blockchain height.
func (self *Blockchain) GetHeight() uint {

	return uint(len(self.Blocks) - 1)
}

// This function adds a block to the blockchain.
func (self *Blockchain) AddBlock(block *block.Block) {

	self.Blocks = append(self.Blocks, *block)
}

// This function removes the last block from the blockchain.
// Only affects the most recent block for error protection.
func (self *Blockchain) RemoveBlock() {

	self.Blocks = append(self.Blocks[:self.GetHeight()], self.Blocks[self.GetHeight()+1:]...)
}

// Replaces a block at a given index with another.
// Will be used if a new valid chain has a higher height,
// allowing the latest block on the saved chain to be replaced with
// the new more valid block.
// Only affects the most recent block for error protection.
func (self *Blockchain) SwitchBlock(b block.Block) {

	self.Blocks[self.GetHeight()] = b
}

// This function gets a block at a specified index.
// Returns the block and true if this was successful.
// If the index is invalid, it will return a empty block and false.
func (b *Blockchain) GetBlock(blockNum uint) (block.Block, bool) {

	if blockNum > b.GetHeight() {

		return block.Block{}, false
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

		unPacker := new(miner.TargetUnpacker)
		packer := new(miner.TargetPacker)
		byteUtil := new(util.ByteUtil)
		time := b.Blocks[blockNumber-1].Timestamp - b.Blocks[blockNumber-10080].Timestamp

		newMultiplier := (10080 * 60) / time // The *60 converts to seconds

		// Convert this to a uint256
		bigNewMultiplier := *types.NewUInt256("0", 1)
		bigNewMultiplier.Set(byteUtil.Uint64toB(newMultiplier))

		// Convert the current target to uint256
		target := unPacker.Unpack(b.Blocks[blockNumber-1].PackedTarget)

		// Apply the multiplier to the current target to get the new target
		newTarget := target.Multiply(&bigNewMultiplier)
		maxTarget := unPacker.Unpack(0x1d0fffff)

		// If the target is larger than the max allowed target
		if newTarget.Compare(&maxTarget) == 1 {

			return 0x1d0fffff
		}

		return packer.PackTargetUint256(*newTarget)
	}

	return b.Blocks[blockNumber-1].PackedTarget
}

// This function saves the blockchain to the computers hard-disk.
// Input is the name of the blockchain being saved.
// Returns nothing.
func (b *Blockchain) SaveBlockchain(bcName string) {

	err := os.WriteFile("saves/"+bcName+".json", b.AsBytes(), 0750)

	if err != nil {

		panic(err)
	}
}

// Loads a saved blockchain.
// Input is the name of the blockchain.
// Returns nothing.
func (b *Blockchain) LoadBlockchain(bcName string) {

	bAsBytes, err := os.ReadFile("saves/" + bcName + ".json")

	if err != nil {

		panic(err)
	}

	// Convert the data to a blockchain from json
	err = json.Unmarshal(bAsBytes, b)

	if err != nil {

		panic(err)
	}
}

// Converts the blockchain into its bytes,
// Returns the byte slice of the blockchain.
func (b *Blockchain) AsBytes() []byte {

	// Get the byte slice
	bAsBytes, err := json.Marshal(b)

	if err != nil {

		panic(err)
	}

	return bAsBytes
}

// This function gets the current difficulty of the blockchain.
// No inputs required and returns the uint64 of the current difficulty.
func (b *Blockchain) GetDifficulty() uint64 {

	unpacker := new(miner.TargetUnpacker)

	currentTarget := unpacker.Unpack(b.Blocks[b.GetHeight()].PackedTarget)
	genisisTarget := unpacker.Unpack(0x1d0fffff)

	difficulty := genisisTarget.Divide(&currentTarget)

	// Returns the uint256 as a uint64 from little endian order
	return binary.LittleEndian.Uint64(difficulty.ToBytes())
}

// This function gets the difficulty of a specific block from the blockchain.
// Only input is the block number and returns the uint64 of that blocks difficulty.
func (b *Blockchain) GetDifficultyOfBlock(blockN uint) uint64 {

	unpacker := new(miner.TargetUnpacker)

	currentTarget := unpacker.Unpack(b.Blocks[blockN].PackedTarget)
	genisisTarget := unpacker.Unpack(0x1d0fffff)

	difficulty := genisisTarget.Divide(&currentTarget)

	// Returns the uint256 as a uint64 from little endian order
	return binary.LittleEndian.Uint64(difficulty.ToBytes())
}
