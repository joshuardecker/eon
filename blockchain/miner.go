package blockchain

import (
	"bytes"
	"encoding/hex"
	"errors"
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

	util     utilities.ByteUtil
	unpacker utilities.TargetUnpacker
	utilTime utilities.Time
}

// Starts the miner. Will return a byte array of the valid hash once discovered. Also returns an error if once occured.
func (m *Miner) Start(b *Block, difficulty uint64) error {

	//****
	// Prepare the miner

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

			// Set the block hash to the winning hash
			b.BlockHash = hex.EncodeToString(m.currentHash)

			fmt.Println("[MINER]:", color.Colorize(color.Green, "Block Found!"))

			m.blocksFound += 1

			return nil
		}

		// Prints stats every 10 MHs
		if b.Nonce%10000000 == 0 {

			timer = m.utilTime.Timer()

			if timer != 0 {

				fmt.Println("!==========!")

				fmt.Println("[MINER]:", color.Colorize(color.Yellow, "Mining..."))
				fmt.Println("[MINER]:", m.utilTime.CurrentTime())
				fmt.Printf("[MINER]: Heres a random of the hashes: %x\n", m.currentHash)
				fmt.Println("[MINER]: Current Difficulty:", difficulty, "| Blocks Found:", m.blocksFound)
				fmt.Println("[MINER]: Average Hashing Speed: ", ((10000000/timer)*60)/1000000, " MH / per minute.")

				fmt.Println("!==========!")
			}
		}

		// Mining
		//****
	}

	return errors.New("you have reached the end of the defined search space! Impressive")
}
