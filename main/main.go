package main

import (
	"flag"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/wallet"
)

// To make a windows copy: GOOS=windows GOARCH=amd64 go build -o luncheon.exe main/main.go
// Website I got this from: https://freshman.tech/snippets/go/cross-compile-go-programs/

func main() {

	//****
	// Input Handler

	test := flag.Bool("testnet", false, "Activates the Testnet")
	flag.Parse()

	if *test {

		fmt.Println("Testnet!")
	}

	// Input Handler
	//****

	fmt.Println("*****")
	//****
	// Test new features area

	// Test new features area
	//****
	fmt.Println("*****")

	//****
	// Starts Mining

	bc := blockchain.InitBlockchain()
	miner := new(blockchain.Miner)
	wal := wallet.Init(&bc)
	key := new(ellip.MainKey)

	miner.Start(&bc.Blocks[0])

	bc.Blocks[0].PrintBlock()

	fmt.Println("Balance:", wal.ScanChainForBalance(key.GetPubKeyStr()))

	// Starts the mining
	//****
}
