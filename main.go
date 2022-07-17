package main

import (
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/luncheon"
)

func main() {

	//****
	// Init the needed structs

	unpacker := new(luncheon.TargetUnpacker)
	packer := new(luncheon.TargetPacker)

	// Init the needed structs
	//****

	//****
	// Section Unpacks the target

	target := unpacker.UnpackAsBytes(0x1d00ffff) // Current btc bits: 0x1709a7af

	// Recalculates the packed target, right now thats for debugging purposes / testing
	packedTarget, packErr := packer.PackTargetBytes(target)

	// Did that func give errors?
	if packErr != nil {

		panic(packErr)
	}

	fmt.Printf("Packed Target: %x\n", 0x1d00ffff)
	fmt.Printf("Recalculated Packed Target: %x\n", packedTarget)

	// Section Unpacks the target
	//****

	//****
	// Starts Mining

	miner := new(luncheon.Miner)

	dataErr := miner.InputData("Hello World!!")

	if dataErr != nil {
		fmt.Println(dataErr)
	}

	targetErr := miner.InputTarget(0x1d00ffff)

	if targetErr != nil {
		fmt.Println(targetErr)
	}

	hash, minerErr := miner.Start()

	if minerErr != nil {
		fmt.Println(minerErr)
	}

	fmt.Printf("Hash: %x\n", hash)

	// Starts the mining
	//****
}
