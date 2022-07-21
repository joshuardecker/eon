package blockchain

import (
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

type Mempool struct {
	Txs []transactions.LuTx
}

// Function adds a tx to the mempool of the blockchain.
// Inputs the transaction you are adding.
// Returns nothing.
func (m *Mempool) AddTx(tx transactions.LuTx) {

	m.Txs = append(m.Txs, tx)
}

// This function removes a tx from the mempool.
// Returns nothing.
func (m *Mempool) RemoveTx(index int) {

	m.Txs = append(m.Txs[:index], m.Txs[index+1:]...)
}

func (m *Mempool) GetTx() any {

	leng := len(m.Txs)

	if leng == 0 {

		return false
	}

	for index := 0; index < leng; index += 1 {

		if m.CheckTxFlags(index) {

			tx := m.Txs[index]
			m.RemoveTx(index)

			return tx
		} else {

			m.RemoveTx(index)

			// So when it loops, it will be the same index value
			index -= 1
		}
	}

	return false
}

// This function checks whether a transaction has the correct flags or not.
// Returns a bool if it is valid, true if yes, false if no.
func (m *Mempool) CheckTxFlags(txIndex int) bool {

	// Save a copy of the tx
	tx := m.Txs[txIndex]

	// Define what we are looking for
	var keyWords []string = []string{
		"TXID",     // Input
		"OUTINDEX", // Input
		"PUBK",     // Output
		"SIG",      // Output
		"PUBKH",    // Output
		"AMT",      // Output
		"SELF",     // Output, but later check
	}

	//****
	// Check the input scripts

	for index := 0; index < len(tx.InScripts); index += 1 {

		// Parse the input scriptstr into a script
		inputScript := transactions.ScriptParse(tx.InScripts[index].ScriptStr)

		// Check keyword, starting at index 0, aka "TXID"
		keyWordIndex := 0

		// Check through the whole scipt
		for nIndex := 0; nIndex < len(inputScript); nIndex += 1 {

			// If the correct word was found
			if inputScript[nIndex] == keyWords[keyWordIndex] {

				keyWordIndex += 1

				// When loops again, it will equal 0
				nIndex = -1
			}
		}

		// Since we were looking for checking for 2 input flags, it must equal 2
		if keyWordIndex != 2 {

			return false
		}
	}

	// Check the input scripts
	//****

	//****
	// Check output scripts

	for index := 0; index < len(tx.OutScripts); index += 1 {

		// Parse the input scriptstr into a script
		outScript := transactions.ScriptParse(tx.OutScripts[index].ScriptStr)

		// Check keyword, starting at index 2, aka "PUBK"
		keyWordIndex := 2

		// Check through the whole scipt
		for nIndex := 0; nIndex < len(outScript); nIndex += 1 {

			// If the correct word was found
			if outScript[nIndex] == keyWords[keyWordIndex] {

				keyWordIndex += 1

				// When loops again, it will equal 0
				nIndex = -1
			}
		}

		// Since we were looking for checking for 4 input flags, it must equal 6 (4 + the starting value 2 = 6)
		if keyWordIndex != 6 {

			return false
		}
	}

	// Check output scripts
	//****

	return true
}
