package main

import (
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
)

func main() {

	fmt.Println("*****")
	//****
	// Test new features area

	mainKey := new(ellip.MainKey)
	mainKey.MainKeyHash()

	// Test new features area
	//****
	fmt.Println("*****")

	//****
	// Starts Mining

	miner := new(blockchain.Miner)

	block := new(blockchain.Block)
	block.InitBlock("Cool Test", 0x1dffffff, 200)

	_, minerErr := miner.Start(*block)

	if minerErr != nil {

		panic(minerErr)
	}

	// Starts the mining
	//****
}
