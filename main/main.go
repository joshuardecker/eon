package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/client"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/mempool"
	"github.com/Sucks-To-Suck/LuncheonNetwork/wallet"
)

func Test(bc *blockchain.Blockchain, mem *mempool.Mempool) {

	miner := new(blockchain.Miner)
	keys := new(ellip.MainKey)

	// Mine the genisis block
	miner.Start(&bc.Blocks[0])

	// Create an endless loop of blockchaining
	for {

		// Create the new block
		block := bc.CreateBlock(keys.GetPubKeyStr())

		// Add the txs
		for txIndex := 0; txIndex < len(mem.Txs); txIndex += 1 {

			tx := mem.GetTx()

			if !block.AddTx(tx) {

				mem.AddTx(&tx)
				break
			}
		}

		// Start mining
		miner.Start(&block)

		// Print the mined block
		block.PrintBlock()

		// Add the newly mined block
		bc.AddBlock(&block)

		// Save the blockchain
		bc.SaveBlockchain("localBlockchain")
	}
}

// To make a windows copy: GOOS=windows GOARCH=amd64 go build -o luncheon.exe main/main.go
// Website I got this from: https://freshman.tech/snippets/go/cross-compile-go-programs/

func main() {

	//****
	// Input Handler

	localServer := flag.Bool("local", false, "Starts a locally hosted testnet")
	localServerTx := flag.Bool("localTx", false, "Sends a tx on the local testnet")

	flag.Parse()

	if *localServer {

		// Init everything needed for the local testnet
		bc := blockchain.InitBlockchain()
		wallet := wallet.Init(&bc)
		mem := mempool.Init(&wallet)

		serverMux := client.Init(&mem)
		mux := serverMux.InitMux()

		// Run the server locally, and as a go routine, the sudo multi threading.
		go http.ListenAndServe(":1919", &mux)

		go Test(&bc, &mem)

		fmt.Scanln()
	}

	if *localServerTx {

		bc := blockchain.InitBlockchain()
		bc.LoadBlockchain("localBlockchain")

		wallet := wallet.Init(&bc)

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
