package wallet

import (
	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

// The version of the software here
var SoftwareVersion string = "v1"

type Wallet struct {
	chain   *blockchain.Blockchain
	version string
}

func Init(b *blockchain.Blockchain) Wallet {

	w := new(Wallet)

	w.chain = b
	w.version = SoftwareVersion

	return *w
}

// Scans the blockchain for the available balance of a publicKey.
// Returns the balance of the publicKey.
func (w *Wallet) ScanChainForBalance(pubKey string) (balance uint64) {

	// Scans the blockchain, starting from the newest block to the first
	for index := 0; index < len(w.chain.Blocks); index += 1 {

		// Check if they got the block reward
		if w.chain.Blocks[index].Miner == pubKey {

			balance += w.chain.GetBlockReward(uint32(index))
		}

		// Check each tx in the block
		for txIndex := 0; txIndex < len(w.chain.Blocks[index].Txs); txIndex += 1 {

			if w.chain.Blocks[index].Txs[txIndex].TxTo == pubKey {

				balance += w.chain.Blocks[index].Txs[txIndex].Value
			}
		}
	}

	return balance
}

// Checks the the tx has enough balance to send a transaction.
// Returns true if yes, and false if no.
func (w *Wallet) CheckTxAmount(tx transactions.LuTx) bool {

	if (tx.Value + tx.Fee) > w.ScanChainForBalance(tx.TxFrom) {

		return false
	}

	return true
}
