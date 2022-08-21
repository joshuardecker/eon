package worker

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/Sucks-To-Suck/LuncheonNetwork/core/events/block"
	"github.com/Sucks-To-Suck/LuncheonNetwork/util"
	"github.com/Sucks-To-Suck/LuncheonNetwork/util/byteUtil"
	"github.com/TwiN/go-color"
	"golang.org/x/crypto/sha3"
)

// The struct that handles the mining. Uses the shake256 varient of sha3 for hashing.
// Here is how the miner handles block hashing. (This is the order of the append list) (adding all the info together)
// SoftwareVersion + PrevBlockHash + MerkleRoot + PackedTarget + Time + Nonce
type Worker struct {
	currentHash    []byte
	unpackedTarget []byte
	nonce          uint32
	mining         bool
	hashCount      uint32

	utilTime util.Time
}

// Creates and returns the new worker.
func New(difficulty *big.Int) *Worker {

	w := new(Worker)

	w.unpackedTarget = difficulty.Bytes()
	w.mining = true

	return w
}

// Starts the miner with the inputted block.
// Will stop if the block is found and added to the blockchain seperatly.
// Returns true if it found the block.
func (self *Worker) Start(b *block.Block) string {

	//****
	// Prepare the miner

	fmt.Println("[MINER]:", color.Colorize(color.Yellow, "New Block!"))

	// Prepare the miner
	//****

	// The actual mining process
	for b.Nonce = 0; b.Nonce <= 0xFFFFFFFF; b.Nonce++ {

		//****
		// Var changes in the process

		// Set the timestamp in the block
		b.Timestamp = self.utilTime.CurrentUnix()

		// Get the block as bytes for mining
		softwareVersion := []byte(util.SoftwareVersion)
		prevBlockHash, _ := hex.DecodeString(b.PrevHash)
		merkleRoot, _ := hex.DecodeString(b.MerkleRoot)
		blockTime := byteUtil.Uint64toB(b.Timestamp)
		packedTargetBytes := byteUtil.Uint32toB(b.PackedTarget)
		nonceBytes := byteUtil.Uint32toB(b.Nonce)

		// Shove them together (into softwareVerion var bc it is first declared)
		softwareVersion = append(softwareVersion, prevBlockHash...)
		softwareVersion = append(softwareVersion, merkleRoot...)
		softwareVersion = append(softwareVersion, packedTargetBytes...)
		softwareVersion = append(softwareVersion, blockTime...)
		softwareVersion = append(softwareVersion, nonceBytes...)

		// Init the size of the hash
		self.currentHash = make([]byte, 32)

		// Var changes in the process
		//****

		//****
		// Mining

		// Hash the data
		sha3.ShakeSum256(self.currentHash, softwareVersion)

		// Was the solution found?
		if bytes.Compare(self.currentHash, self.unpackedTarget) != 1 {

			fmt.Println("[MINER]:", color.Colorize(color.Green, "Block Found!"))

			// Set the block hash to the winning hash
			return hex.EncodeToString(self.currentHash)
		}

		// Mining
		//****
	}

	return ""
}
