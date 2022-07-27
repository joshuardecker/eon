package node

import (
	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/mempool"
)

// This function handles the mining and overall blockchain process on a local network.
// Inputs are the blockchain and the mempool of the local node, a miner for it, and the keys of the local node.
// Returns nothing.
func LocalNodeMining(bc *blockchain.Blockchain, mem *mempool.Mempool, miner *blockchain.Miner, keys *ellip.MainKey) {

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
