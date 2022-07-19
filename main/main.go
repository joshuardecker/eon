package main

import (
	"flag"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

func main() {

	//****
	// Input Handler

	test := flag.Bool("testnet", false, "activates the testnet")
	flag.Parse()

	if *test {

		fmt.Println("Testnet!")
	}

	// Input Handler
	//****

	fmt.Println("*****")
	//****
	// Test new features area

	blockChain := new(blockchain.Blockchain)
	block := new(blockchain.Block)
	plux := new(transactions.PLuX)

	block.InitBlock("8f2348098a", 0x1dffffff, blockChain.GetBlockReward())

	plux.CreatePLuX(blockChain.GetBlockReward())

	block.AddPLuX(*plux)

	fmt.Printf("Packed target: %x\n", block.PackedTarget)

	// Test new features area
	//****
	fmt.Println("*****")

	//****
	// Starts Mining

	miner := new(blockchain.Miner)

	finalBlock, minerErr := miner.Start(*block)

	if minerErr != nil {

		panic(minerErr)
	}

	// Prints the block as a json string
	finalBlock.PrintBlock()

	blockChain.AddBlock(finalBlock)

	// Starts the mining
	//****
}
