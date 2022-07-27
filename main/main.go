package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/mempool"
	"github.com/Sucks-To-Suck/LuncheonNetwork/node"
	"github.com/Sucks-To-Suck/LuncheonNetwork/wallet"
)

// To make a windows copy: GOOS=windows GOARCH=amd64 go build -o luncheon.exe main/main.go
// Website I got this from: https://freshman.tech/snippets/go/cross-compile-go-programs/

func main() {

	//****
	// Input Handler

	localNode := flag.Bool("local", false, "Starts a locally hosted testnet")
	localNodeTx := flag.Bool("localTx", false, "Sends a tx on the local testnet")

	flag.Parse()

	// Init vars needed for the blockchain processes below
	bc := blockchain.InitBlockchain()
	wallet := wallet.Init(&bc)
	mem := mempool.Init(&wallet)
	miner := new(blockchain.Miner)
	keys := new(ellip.MainKey)

	if *localNode {

		serverMux := node.Init(&mem)
		mux := serverMux.InitMux()

		// Run the server locally, and as a go routine, the sudo multi threading.
		go http.ListenAndServe(":1919", &mux)

		go node.LocalNodeMining(&bc, &mem, miner, keys)

		bc.GetDifficulty()

		fmt.Scanln()
	}

	if *localNodeTx {

		tx := wallet.CreateTx("kaimort123", 10)
		txBuffer := bytes.NewBuffer(tx.AsBytes())

		_, err := http.Post("http://localhost:1919/tx", "tx/text", txBuffer)

		if err != nil {

			panic(err)
		}
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

	// Starts the mining
	//****
}
