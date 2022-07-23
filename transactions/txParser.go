package transactions

import (
	"encoding/hex"
	"strconv"
)

// This function takes in a tx input (assummed to be verified) and gives
// all of the data related to the tx.
// Returns the blocknumbers referenced in the input script, there tx hashes,
// and there output indexes.
// Next returns the public key in the tx, the signature, the hash of the public key
// where the coin is going, and the amount of coin going there.
func TransactionParser(tx LuTx) (valid bool, blockNumber []uint64, txId []string,
	pubKey []byte, signature []byte, pubKeyHash string, amount uint64, nTxId string) {

	//****
	// Get the InputScript data

	// Loops through all of the inputScripts
	for index := 0; index < len(tx.InScripts); index += 1 {

		script := StrToScript(tx.InScripts[index].ScriptStr)

		// Loops through the script
		for nIndex := 0; nIndex < len(script); nIndex += 1 {

			// Block reward
			if script[nIndex] == "BLKN" {

				bN, err := strconv.ParseUint(script[nIndex+1], 10, 32)

				if err != nil {

					return false, blockNumber, txId, pubKey, signature, pubKeyHash, amount, nTxId
				}

				blockNumber = append(blockNumber, bN)

				nIndex += 1
			}

			// Transaction id hash
			if script[nIndex] == "TXID" {

				if len(script[nIndex+1]) != 64 {

					return false, blockNumber, txId, pubKey, signature, pubKeyHash, amount, nTxId
				}

				txId = append(txId, script[nIndex+1])

				nIndex += 1
			}
		}
	}

	// Get the InputScript data
	//****

	//****
	// Get the OutputScript data

	script := StrToScript(tx.OutScripts.ScriptStr)

	// Loops through the script
	for nIndex := 0; nIndex < len(script); nIndex += 1 {

		if script[nIndex] == "PUBK" {

			pK, err := hex.DecodeString(script[nIndex+1])

			if err != nil {

				return false, blockNumber, txId, pubKey, signature, pubKeyHash, amount, nTxId
			}

			pubKey = pK

			nIndex += 1
		}

		if script[nIndex] == "SIG" {

			sG, err := hex.DecodeString(script[nIndex+1])

			if err != nil {

				return false, blockNumber, txId, pubKey, signature, pubKeyHash, amount, nTxId
			}

			signature = sG

			nIndex += 1
		}

		if script[nIndex] == "PUBKH" {

			if len(script[nIndex+1]) != 64 {

				return false, blockNumber, txId, pubKey, signature, pubKeyHash, amount, nTxId
			}

			pubKeyHash = script[nIndex+1]

			nIndex += 1
		}

		if script[nIndex] == "AMT" {

			aT, err := strconv.ParseUint(script[nIndex+1], 10, 64)

			if err != nil {

				return false, blockNumber, txId, pubKey, signature, pubKeyHash, amount, nTxId
			}

			amount = aT

			nIndex += 1
		}

		if script[nIndex] == "NTXID" {

			if len(script[nIndex+1]) != 64 {

				return false, blockNumber, txId, pubKey, signature, pubKeyHash, amount, nTxId
			}

			nTxId = script[nIndex+1]

			nIndex += 1
		}
	}

	// Get the OutputScript data
	//****

	return true, blockNumber, txId, pubKey, signature, pubKeyHash, amount, nTxId
}
