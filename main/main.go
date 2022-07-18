package main

import (
	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
)

func main() {

	//****
	// Starts Mining

	miner := new(blockchain.Miner)

	block := new(blockchain.Block)
	block.ConstructBlock()

	targetErr := miner.InputTarget(0x1dffffff) //1d00ffff

	if targetErr != nil {

		panic(targetErr)
	}

	_, minerErr := miner.Start(*block)

	if minerErr != nil {

		panic(minerErr)
	}

	// Starts the mining
	//****
}
