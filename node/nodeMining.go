package node

import (
	"bytes"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/crypto/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/mempool"
	"github.com/Sucks-To-Suck/LuncheonNetwork/wallet"
	"github.com/TwiN/go-color"
)

type NodeMiner struct {
	node   *Node
	bc     *blockchain.Blockchain
	mem    *mempool.Mempool
	miner  *blockchain.Miner
	keys   *ellip.MainKey
	wallet *wallet.Wallet

	saveName string
}

// This function creates a node miner with specified inputs.
// The saveName input is what the save will be named when the blockchain is saved to the harddisk.
// Returns a pointer to the NodeMiner.
func InitNodeMiner(node *Node, bc *blockchain.Blockchain, mem *mempool.Mempool, miner *blockchain.Miner, keys *ellip.MainKey, wallet *wallet.Wallet,
	saveName string) *NodeMiner {

	nm := new(NodeMiner)

	// Set vars
	nm.node = node
	nm.bc = bc
	nm.mem = mem
	nm.miner = miner
	nm.keys = keys
	nm.saveName = saveName
	nm.wallet = wallet

	return nm
}

func (nm *NodeMiner) StartMining() {

	// Save the empty blockchain
	nm.bc.SaveBlockchain(nm.saveName)

	// Mine the genisis block
	nm.miner.Start(&nm.bc.Blocks[0], nm.bc, nm.bc.GetDifficulty())

	// Save the genisis block
	nm.bc.SaveBlockchain(nm.saveName)

	// Create an endless loop of blockchaining
	for {

		// Create the new block
		block := nm.bc.CreateBlock(nm.keys.GetPubKeyStr())

		// Add the block to the chain
		// It needs to be added now so wallet functions can check if there
		// Are duplicate txs in a block, as without it, multiple copys of the same
		// valid tx could be put in the block, making this node waste time mining an invalid block.
		nm.bc.AddBlock(&block)

		// Add the txs
		for txIndex := 0; txIndex < len(nm.mem.Txs); txIndex += 1 {

			tx := nm.mem.GetTx()

			// If the tx is valid, add it
			if nm.wallet.VerifyTx(tx) {

				// If it could not be added, put it back in the mempool
				if !block.AddTx(tx) {

					nm.mem.AddTx(&tx)
					break
				}
			}
		}

		// If the miner finds the block
		if nm.miner.Start(&block, nm.bc, nm.bc.GetDifficulty()) {

			// Add the newly mined block
			nm.bc.AddBlock(&block)

			// Save the blockchain
			nm.bc.SaveBlockchain(nm.saveName)

			buffer := bytes.NewBuffer(block.AsBytes())

			// Tell all known peers about the block
			nm.node.SendDataToAll("/newblock", buffer)

			fmt.Println(color.Colorize(color.Green, "[NODE]: Successfully sent block to peers"))
		}
	}
}
