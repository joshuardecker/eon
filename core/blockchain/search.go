package blockchain

// Scans the blockchain for the available balance of a publicKey.
// Returns the balance of the publicKey.
func (self *Blockchain) ScanChainForBalance(pubKey string) (balance uint64) {

	// Scans the blockchain, starting from the newest block to the first
	for index := 0; index < len(self.Blocks); index += 1 {

		// Check if they got the block reward (+10 makes the miner wait at least 10 blocks before it can be spent)
		if self.Blocks[index].Miner == pubKey && (index+10) < int(self.GetHeight()) {

			balance += self.GetBlockReward(uint32(index))
		}

		// Check each tx in the block
		for txIndex := 0; txIndex < len(self.Blocks[index].Txs); txIndex += 1 {

			// If coin received
			if self.Blocks[index].Txs[txIndex].TxTo == pubKey {

				balance += self.Blocks[index].Txs[txIndex].Value
			}

			// If coin spent
			if self.Blocks[index].Txs[txIndex].TxFrom == pubKey {

				balance -= self.Blocks[index].Txs[txIndex].Value
			}
		}
	}

	return balance
}

// Scans the blockchain for the available balance of a publicKey.
// Returns the balance of the publicKey.
func (self *Blockchain) ScanChainForNonce(pubKey string) (nonce uint32) {

	// Scans the blockchain, starting from the newest block to the first
	for index := 0; index < len(self.Blocks); index += 1 {

		// Check each tx in the block
		for txIndex := 0; txIndex < len(self.Blocks[index].Txs); txIndex += 1 {

			if self.Blocks[index].Txs[txIndex].TxFrom == pubKey {

				nonce += 1
			}
		}
	}

	return nonce
}
