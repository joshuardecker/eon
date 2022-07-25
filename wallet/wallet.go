package wallet

import (
	"encoding/hex"
	"encoding/json"

	"github.com/Sucks-To-Suck/LuncheonNetwork/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
	"golang.org/x/crypto/sha3"
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

// This function creates a tx and verifys it.
// Inputs are the publicKey the tx is going to, and the amount of Luncheon that is being sent.
// Outputs are the tx, which if empty, means that the amount specified is not possible with your balance.
func (w *Wallet) CreateTx(toPub string, amount uint64) (tx transactions.LuTx) {

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

	return tx
}

// Function calculates whether the tx input is valid or not.
// Input is the tx.
// Returns true if valid, false if not valid.
func (w *Wallet) VerifyTx(tx transactions.LuTx) bool {

	// If the tx has a spendable amount of coin from the persons balance
	if (tx.Value + tx.Fee) > w.ScanChainForBalance(tx.TxFrom) {

		return false
	}

	// If the tx has the wrong nonce value
	if tx.Nonce != w.ScanChainForNonce(tx.TxFrom) {

		return false
	}

	// Remove the sig from the tx and save it, as to get the tx hash input data
	signature, _ := hex.DecodeString(tx.Signature)
	tx.Signature = ""

	txBytes := tx.AsBytes()
	txHash := make([]byte, 32)
	pubKey, _ := hex.DecodeString(tx.TxFrom)

	sha3.ShakeSum256(txHash, txBytes)

	// If the signature is not valid
	// If this is true, than the tx is true
	return ellip.ValidateSig(pubKey, txHash, signature)
}
