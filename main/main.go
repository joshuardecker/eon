package main

import (
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
)

func main() {

	pub, priv := ellip.GetKeyPair()
	fmt.Printf("%x\n", pub)
	fmt.Printf("%x\n", priv)

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
