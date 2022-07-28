package blockchain

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/utilities"
	"github.com/TwiN/go-color"
	"golang.org/x/crypto/sha3"
)

// The struct that handles the mining. Uses the shake256 varient of sha3 for hashing.
// Here is how the miner handles block hashing. (This is the order of the append list) (adding all the info together)
// SoftwareVersion + PrevBlockHash + MerkleRoot + PackedTarget + Time + Nonce
type Miner struct {
	currentHash    []byte
	unpackedTarget []byte
	blocksFound    uint
	startHeight    uint

	util     utilities.ByteUtil
	unpacker utilities.TargetUnpacker
	utilTime utilities.Time
}

// Starts the miner with the inputted block.
// Will stop if the block is found and added to the blockchain seperatly.
// Returns true if it found the block.
func (m *Miner) Start(b *Block, bc *Blockchain, difficulty uint64) bool {

	//****
	// Prepare the miner

	m.startHeight = bc.GetHeight()

	// Gets the unpacked target with the unpacker struct
	m.unpackedTarget = m.unpacker.UnpackAsBytes(b.PackedTarget)

	// Init the timer used for calculating MH/s
	timer := m.utilTime.Timer()

	fmt.Println("[MINER]:", color.Colorize(color.Yellow, "New Block!"))

	// Prepare the miner
	//****

	// The actual mining process
	for b.Nonce = 0; b.Nonce <= 0xFFFFFFFF; b.Nonce++ {

		//****
		// Var changes in the process

		// Set the timestamp in the block
		b.Timestamp = m.utilTime.CurrentUnix()

		// Get the block as bytes for mining
		softwareVersion := []byte(utilities.SoftwareVersion)
		prevBlockHash, _ := hex.DecodeString(b.PrevHash)
		merkleRoot, _ := hex.DecodeString(b.MerkleRoot)
		blockTime := m.util.Uint64toB(b.Timestamp)
		packedTargetBytes := m.util.Uint32toB(b.PackedTarget)
		nonceBytes := m.util.Uint32toB(b.Nonce)

		// Shove them together (into softwareVerion var bc it is first declared)
		softwareVersion = append(softwareVersion, prevBlockHash...)
		softwareVersion = append(softwareVersion, merkleRoot...)
		softwareVersion = append(softwareVersion, packedTargetBytes...)
		softwareVersion = append(softwareVersion, blockTime...)
		softwareVersion = append(softwareVersion, nonceBytes...)

		// Init the size of the hash
		m.currentHash = make([]byte, 32)

		// Var changes in the process
		//****

		//****
		// Mining

		// Hash the data
		sha3.ShakeSum256(m.currentHash, softwareVersion)

		// Was the solution found?
		if bytes.Compare(m.currentHash, m.unpackedTarget) != 1 {

			// Check if the block has been found
			if m.startHeight != bc.GetHeight() {

				fmt.Println("[MINER]:", color.Colorize(color.Yellow, "Found solution to old block. Scrapping old block..."))
				return false
			}

			// Set the block hash to the winning hash
			b.BlockHash = hex.EncodeToString(m.currentHash)

			fmt.Println("[MINER]:", color.Colorize(color.Green, "Block Found!"))

			m.blocksFound += 1

			return true
		}

		// Prints stats every 20 MHs
		if b.Nonce%20000000 == 0 {

			// Check if the block has been found
			if m.startHeight != bc.GetHeight() {

				fmt.Println("[MINER]:", color.Colorize(color.Yellow, "Scrapping old block..."))
				return false
			}

			timer = m.utilTime.Timer()

			if timer != 0 {

				fmt.Println("!==========!")

				fmt.Println("[MINER]:", color.Colorize(color.Yellow, "Mining..."))
				fmt.Println("[MINER]:", m.utilTime.CurrentTime())
				fmt.Printf("[MINER]: Heres a random of the hashes: %x\n", m.currentHash)
				fmt.Println("[MINER]: Current Difficulty:", difficulty, "| Blocks Found:", m.blocksFound)
				fmt.Println("[MINER]: Average Hashing Speed: ", ((20000000/timer)*60)/1000000, " MH / per minute.")

				fmt.Println("!==========!")
			}
		}

		// Mining
		//****
	}

	return false
}
