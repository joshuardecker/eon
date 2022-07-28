package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/mempool"
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
	"github.com/Sucks-To-Suck/LuncheonNetwork/wallet"
	"github.com/TwiN/go-color"
)

type Node struct {
	bc      *blockchain.Blockchain
	mem     *mempool.Mempool
	wal     *wallet.Wallet
	mainnet bool

	Peers []string
}

// Inits the Node.
// Input is the mempool that will be used by the Node and whether this node is on the mainnet. (True if yes).
// Returns the new Node.
func Init(bc *blockchain.Blockchain, mempool *mempool.Mempool, mainnet bool, wallet *wallet.Wallet) *Node {

	n := new(Node)

	// Save inputted vars
	n.bc = bc
	n.mem = mempool
	n.mainnet = mainnet
	n.wal = wallet

	return n
}

// Initiates the mux for the server.
// Returns the ServerMux of all of the Handled functions of the client.
func (n *Node) InitMux() http.ServeMux {

	mux := http.NewServeMux()

	// Handled Funcs
	mux.HandleFunc("/tx", n.AddTx)
	mux.HandleFunc("/status", n.Status)
	mux.HandleFunc("/newblock", n.Newblock)
	mux.HandleFunc("/getbc", n.SendBlockchain)

	return *mux
}

// Adds a tx to the mempool.
// Inputs are the requests and writer from the http request.
// Returns nothing.
// Accessed by "/tx".
func (n *Node) AddTx(w http.ResponseWriter, r *http.Request) {

	// Get the body of the http message.
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Could not read transaction. Err:") + err.Error())

		// Tells the client that the tx was not accepted
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	tx := new(transactions.LuTx)

	// Get the tx into the tx struct
	err = json.Unmarshal(body, tx)

	if err != nil {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Could not unmarshal transaction. Err:") + err.Error())

		// Tells the client that the tx was not accepted
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// Add the tx to the mempool
	n.mem.AddTx(tx)

	// Tells the client that the tx was accepted
	w.WriteHeader(http.StatusAccepted)

	fmt.Println(color.Colorize(color.Green, "[NODE]: Successfully received new transaction."))

	// Send the tx to all your known peers
	buffer := bytes.NewBuffer(body)
	n.SendDataToAll("/tx", buffer)
}

// This function simply allows to see if each other are online, and you to them the same.
// This is used mainly in p2p, where nodes get each others ips.
// Returns nothing.
func (n *Node) Status(w http.ResponseWriter, r *http.Request) {

	// Give a valid response
	w.WriteHeader(http.StatusOK)

	// Attempt to add the node who contacted you
	n.AddNode(r.RemoteAddr, n.mainnet)
}

// This allows nodes to share new blocks with each other.
// It will verify a block before adding the block to the chain.
// If valid, it will be added to the chain and the miner will move to the next block.
// Returns nothing.
func (n *Node) Newblock(w http.ResponseWriter, r *http.Request) {

	// Get the body of the http message.
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Could not read sent block. Err:") + err.Error())

		// Tells the client that the tx was not accepted
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	block := new(blockchain.Block)

	err = json.Unmarshal(body, block)

	if err != nil {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Could not unmarshal sent block. Err:") + err.Error())

		// Tells the client that the tx was not accepted
		w.WriteHeader(http.StatusNotAcceptable)

		return
	}

	// If the block is not valid
	if !n.wal.VerifyBlock(block, true) {

		fmt.Println(color.Colorize(color.Red, "[NODE]: Invalid block sent by peer."))

		w.WriteHeader(http.StatusNotAcceptable)

		return
	}

	// Add the block to the chain
	n.bc.AddBlock(block)

	// Send an all good back to the node
	w.WriteHeader(http.StatusAccepted)

	fmt.Println(color.Colorize(color.Green, "[NODE]: Successfully received new valid block."))
}

// Sends the blockchain to a user who requests it.
// Used when nodes are syncing to the network.
// Returns nothing.
func (n *Node) SendBlockchain(w http.ResponseWriter, r *http.Request) {

	// Runs the blockchain writting to a seperate go-routine
	w.Write(n.bc.AsBytes())

	w.WriteHeader(http.StatusOK)
}
