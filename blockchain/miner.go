package blockchain

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/utilities"
	"golang.org/x/crypto/sha3"
)

// The struct that handles the mining. Uses the shake256 varient of sha3 for hashing.
type Miner struct {
	inputBlockBytes []byte
	packedTarget    uint32

	hashData       []byte
	currentHash    []byte
	unpackedTarget []byte

	util     utilities.ByteUtil
	unpacker utilities.TargetUnpacker
}

// This function tells the miner what target to mine to. Returns an error if once occurs.
func (m *Miner) InputTarget(inputTarget uint32) error {

	// 0 is an invalid target, and this handles that
	if inputTarget == 0 {

		return errors.New("cannot input target 0")
	}

	// Is the target higher than the max allowed target?
	if inputTarget > 0x1dffffff { // TODO: Find a better value and define as const elsewhere

		return errors.New("target is to large") // TODO: Have it print max target
	}

	m.packedTarget = inputTarget

	return nil
}

// Starts the miner. Will return a byte array of the valid hash once discovered. Also returns an error if once occured.
func (m *Miner) Start(b Block) ([]byte, error) {

	// Get the block as bytes for mining
	m.inputBlockBytes = b.ParseBlockToBytes()

	// No block data?
	if m.inputBlockBytes == nil {

		return nil, errors.New("please input a block with data inside it")
	}

	// 0 is an invalid target, and this handles that
	if m.packedTarget == 0 {

		return nil, errors.New("please define a non 0 target with the InputTarget function")
	}

	// Gets the unpacked target with the unpacker struct
	m.unpackedTarget = m.unpacker.UnpackAsBytes(m.packedTarget)

	fmt.Println("Mining Starting!")

	// The actual mining process
	for b.Nonce = 0; b.Nonce <= 0xFFFFFFFF; b.Nonce++ {

		// Create the input bytes for the hash, and add the nonce
		m.hashData = append(m.inputBlockBytes, m.util.Uint32toB(b.Nonce)...)

		// Init the size of the hash
		m.currentHash = make([]byte, 32)

		// Hash the data
		sha3.ShakeSum256(m.currentHash, m.hashData)

		// Was the solution found?
		if bytes.Compare(m.currentHash, m.unpackedTarget) != 1 {

			// Set the block hash to the winning hash
			b.SetBlockHash(m.currentHash)

			// Prints the block as a json string
			b.PrintBlock()

			return m.currentHash, nil
		}

		// Prints stats every 10 MH
		if b.Nonce%10000000 == 0 {

			fmt.Println("Mining...")
			fmt.Printf("Target: %x\n", m.unpackedTarget)
			fmt.Printf("Last Hash: %x\n", m.currentHash)
		}
	}

	return nil, errors.New("you have reached the end of the defined search space! Impressive")
}
