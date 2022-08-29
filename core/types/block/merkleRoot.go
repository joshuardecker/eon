package block

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// Function takes all of the transactions in the block,
// and gets their merkle root.
// Returns the hash string of the merkle root.
func (b *Block) GetMerkleRoot() string {

	// Init the length
	leng := len(b.Txs)

	// If there are no txs
	if leng == 0 {

		return ""
	}

	// Save a copy of the tx's
	txCopy := b.Txs

	// Is the total amount of tx's odd?
	if leng%2 == 1 {

		// Makes the tranaction amount even
		txCopy = append(txCopy, txCopy[leng-1])

		// Adjusts the length to be accurate
		leng += 1
	}

	hashStrs := make([]string, leng)

	// Move all of the transactions bytes into a hex encoded string array
	for index := 0; index < leng; index += 1 {

		hashStrs[index] = hex.EncodeToString(txCopy[index].AsBytes())
	}

	// Gets the merkle root
	for index := 0; leng != 1; index += 2 {

		// Init the hash
		hash := make([]byte, 32)

		// Prepare selected the data for hashing (the pairs ex: 0 and 1, 2 and 3)
		byteData, err := hex.DecodeString(hashStrs[index] + hashStrs[index+1])

		// Was there an error?
		if err != nil {
			panic(err)
		}

		// Do the hash
		sha3.ShakeSum256(hash, byteData)

		// Sets the data on the lower end indexes of the array, ex: 0,1,2,3 becomes 0,1
		hashStrs[index/2] = hex.EncodeToString(hash)

		// Stops the index from going out of range
		if (index + 2) == leng {

			// Sets it to -2 because when it loops again, -2 + 2 = 0 and it needs to be set back to 0
			index = -2

			// Half the length because the total data has been halved
			leng /= 2
		}
	}

	return hashStrs[0]
}
