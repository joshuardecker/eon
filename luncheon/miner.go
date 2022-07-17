package luncheon

import (
	"bytes"
	"errors"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// The struct that handles the mining. Uses the shake256 varient of sha3 for hashing.
type Miner struct {
	inputData    string
	packedTarget uint32

	nonce          uint32
	hashData       []byte
	currentHash    []byte
	unpackedTarget []byte

	util     Util
	unpacker TargetUnpacker
}

// This function inputs the string data that will be mined. Returns an error if that occurs.
func (m *Miner) InputData(inputData string) error {

	// Is the string empty?
	if len(inputData) == 0 {

		return errors.New("Use the InputData function!")
	}

	m.inputData = inputData

	return nil
}

// This function tells the miner what target to mine to. Returns an error if once occurs.
func (m *Miner) InputTarget(inputTarget uint32) error {

	// 0 is an invalid target, and this handles that
	if inputTarget == 0 {

		return errors.New("Cannot input target 0!")
	}

	// Is the target higher than the max allowed target?
	if inputTarget > 0x1dffffff { // TODO: Find a better value and define as const elsewhere

		return errors.New("Target is to large!") // TODO: Have it print max target
	}

	m.packedTarget = inputTarget

	return nil
}

// Starts the miner. Will return a byte array of the valid hash once discovered. Also returns an error if once occured.
func (m *Miner) Start() ([]byte, error) {

	// Is the string empty / not set?
	if len(m.inputData) == 0 {

		return nil, errors.New("Use the InputData function!")
	}

	// 0 is an invalid target, and this handles that
	if m.packedTarget == 0 {

		return nil, errors.New("Please define a non 0 target with the InputTarget function!")
	}

	// Gets the unpacked target with the unpacker struct
	m.unpackedTarget = m.unpacker.UnpackAsBytes(m.packedTarget)

	fmt.Println("Mining Starting!")

	// The actual mining process
	for m.nonce = 0; m.nonce <= 0xFFFFFFFF; m.nonce++ {

		// Create the input bytes for the hash, and add the nonce
		m.hashData = append([]byte(m.inputData)[:], m.util.Uint32toB(m.nonce)...)

		// Init the size of the hash
		m.currentHash = make([]byte, 32)

		// Hash the data
		sha3.ShakeSum256(m.currentHash, m.hashData)

		// Was the solution found?
		if bytes.Compare(m.currentHash, m.unpackedTarget) != 1 {

			fmt.Printf("Hash Found!: %x\n", m.currentHash)
			fmt.Println("Nonce: ", m.nonce)

			return m.currentHash, nil
		}

		// Prints stats every 50 MH
		if m.nonce%50000000 == 0 {

			fmt.Println("Mining...")
			fmt.Printf("Target: %x\n", m.unpackedTarget)
			fmt.Printf("Last Hash: %x\n", m.currentHash)
		}
	}

	return nil, errors.New("You have reached the end of the defined search space! Impressive!")
}
