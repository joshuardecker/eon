package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"

	"github.com/Sucks-To-Suck/LuncheonNetwork/core/wallet"
	"github.com/Sucks-To-Suck/LuncheonNetwork/crypto/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/node"
	"github.com/TwiN/go-color"
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

		localNode := node.Init(&bc, &mem, false, &wallet)
		mux := localNode.InitMux()

		// Run the server locally, and as a go routine, the sudo multi threading.
		go http.ListenAndServe(":8180", &mux)

		// Also start the node mining process.
		nodeMiner := node.InitNodeMiner(localNode, &bc, &mem, miner, keys, &wallet, "local")
		go nodeMiner.StartMining()

		// Used to make the program not close, so it waits for user input to stop
		fmt.Scanln()

	} else if *localNodeTx {

		fmt.Println("!==========!")

		// Load the local blockchain
		bc.LoadBlockchain("local")

		// Get and print the available balance
		balance := wallet.ScanChainForBalance(keys.GetPubKeyStr())
		fmt.Println("Available Balance:", balance/1000000)

		if balance == 0 {

			fmt.Println("!==========!")
			return
		}

		fmt.Println("Where is this transaction going?")

		var userToKey string
		_, err := fmt.Scanln(&userToKey)

		// Was there an err getting the users input
		if err != nil {

			fmt.Println(color.Colorize(color.Red, "[TRANSACTION]: Error:"+err.Error()))
			return
		}

		fmt.Println("How much is being sent?")

		var userAmount uint64
		_, err = fmt.Scanln(&userAmount)

		// Was there an err getting the users input
		if err != nil {

			fmt.Println(color.Colorize(color.Red, "[TRANSACTION]: Error:"+err.Error()))
			return

		}

		if (userAmount * 1000000) > balance {

			fmt.Println(color.Colorize(color.Red, "[TRANSACTION]: Error: cannot send more than your balance"))
			return
		}

		tx := wallet.CreateTx(userToKey, userAmount*1000000)

		txBuffer := bytes.NewBuffer(tx.AsBytes())

		// Contact the node
		resp, httpErr := http.Post("http://localhost:8180/tx", "data/json", txBuffer)

		// Was there an error in contacting the node
		if httpErr != nil {

			fmt.Println(color.Colorize(color.Red, "[TRANSACTION]: Error: "+httpErr.Error()))
			return
		}

		// If the node did not respond with the accepted status
		if resp.StatusCode != http.StatusAccepted {

			fmt.Println(color.Colorize(color.Red, "[TRANSACTION]: Error: transaction not accepted by node. Http err"+resp.Status))
			return
		}

		fmt.Println(color.Colorize(color.Green, "[TRANSACTION]: Transaction Accepted by Node!"))

		fmt.Println("!==========!")
	} else {

		fmt.Println("Type 'luncheon -help' to get a list of the possible commands.")
	}

	// Input Handler
	//****
}
