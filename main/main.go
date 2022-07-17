package main

import (
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
)

func main() {

	//****
	// Starts Mining

	miner := new(blockchain.Miner)

	block := new(blockchain.Block)
	block.ConstructBlock()

	targetErr := miner.InputTarget(0x1d00ffff) //1d00ffff

	if targetErr != nil {
		panic(targetErr)
	}

	hash, minerErr := miner.Start(block.ParseBlockToBytes())

	if minerErr != nil {
		panic(minerErr)
	}

	fmt.Printf("Hash: %x\n", hash)

	// Starts the mining
	//****
}
