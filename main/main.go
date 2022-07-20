package main

import (
	"flag"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/utilities"
)

// To make a windows copy: GOOS=windows GOARCH=amd64 go build -o luncheon.exe main/main.go
// Website I got this from: https://freshman.tech/snippets/go/cross-compile-go-programs/

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

	timeUtil := new(utilities.Time)
	fmt.Println(timeUtil.CurrentTime())

	// Test new features area
	//****
	fmt.Println("*****")

	//****
	// Block init

	blockChain := new(blockchain.Blockchain)
	block := new(blockchain.Block)

	block.InitBlock("8f2348098a", 0x1d0fffff) // Right now 1d0fffff is pretty balanced starting target

	// Block init
	//****

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

	fmt.Println(timeUtil.CurrentTime())
}
