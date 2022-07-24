package wallet

import (
	"encoding/hex"
	"encoding/json"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

// The version of the software here
var SoftwareVersion string = "v1"

type Wallet struct {
	chain   *blockchain.Blockchain
	version string
	mainKey ellip.MainKey
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

// Scans the blockchain for the available balance of a publicKey.
// Returns the balance of the publicKey.
func (w *Wallet) ScanChainForNonce(pubKey string) (nonce uint32) {

	// Scans the blockchain, starting from the newest block to the first
	for index := 0; index < len(w.chain.Blocks); index += 1 {

		// Check each tx in the block
		for txIndex := 0; txIndex < len(w.chain.Blocks[index].Txs); txIndex += 1 {

			if w.chain.Blocks[index].Txs[txIndex].TxFrom == pubKey {

				nonce += 1
			}
		}
	}

	return nonce
}

// Checks the the tx has enough balance to send a transaction.
// Returns true if yes, and false if no.
func (w *Wallet) CheckTxAmount(tx transactions.LuTx) bool {

	if (tx.Value + tx.Fee) > w.ScanChainForBalance(tx.TxFrom) {

		return false
	}

	return true
}

// This function creates a tx and verifys it.
// Inputs are the publicKey the tx is going to, and the amount of Luncheon that is being sent.
// Outputs are the tx, which if empty, means that the amount specified is not possible with your balance.
func (w *Wallet) CreateTx(toPub string, amount uint64) transactions.LuTx {

	tx := new(transactions.LuTx)

	// Say the tx is from you
	tx.TxFrom = w.mainKey.GetPubKeyStr()

	tx.TxTo = toPub
	tx.Value = amount

	tx.Nonce = w.ScanChainForNonce(tx.TxFrom)

	// Simple calculation to get a tx fee
	tx.Fee = uint64((tx.GetWeight() + 64) * 1000) // The +64 is to add the weight of the signature

	txBytes, _ := json.Marshal(tx)

	_, sig := w.mainKey.SignMsg(txBytes)
	tx.Signature = hex.EncodeToString(sig)

	return *tx
}
