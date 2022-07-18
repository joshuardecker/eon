package main

import (
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
)

func main() {

	fmt.Println("*****")
	//****
	// Test new features area

	bl := new(blockchain.Blockchain)
	fmt.Println(bl.GetBlockReward())

	// Test new features area
	//****
	fmt.Println("*****")

	//****
	// Starts Mining

	miner := new(blockchain.Miner)

	block := new(blockchain.Block)
	//block.CreateBlock()

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
