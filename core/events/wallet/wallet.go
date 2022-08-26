package wallet

import (
	"encoding/hex"
	"fmt"

	"github.com/Sucks-To-Suck/LuncheonNetwork/core/blockchain"
	"github.com/Sucks-To-Suck/LuncheonNetwork/core/events/block"
	"github.com/Sucks-To-Suck/LuncheonNetwork/core/events/txs"
	"github.com/Sucks-To-Suck/LuncheonNetwork/crypto/ellip"
	"github.com/Sucks-To-Suck/LuncheonNetwork/util"
	"golang.org/x/crypto/sha3"
)

type Wallet struct {
	chain   *blockchain.Blockchain
	mainKey ellip.MainKey
}

// Initialize a wallet by calling this function.
// Input is the blockchain the wallet is on.
// Returns a new wallet.
func Init(b *blockchain.Blockchain) *Wallet {

	w := new(Wallet)

	w.chain = b

	return w
}

// Function calculates whether the tx input is valid or not.
// Input is the tx.
// Returns true if valid, false if not valid.
func (self *Wallet) VerifyTx(tx txs.LuTx) bool {

	// If the tx has a spendable amount of coin from the persons balance
	if int(self.chain.ScanChainForBalance(tx.TxFrom))-int((tx.Value+tx.Fee)) < 0 {

		return false
	}

	// If the tx has the wrong nonce value
	if tx.Nonce != self.chain.ScanChainForNonce(tx.TxFrom) {

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

// Verifies of the block inputted is valid or not.
// Input is the block being verified. The second input is a bool that determines whether a block should have the same software version as you.
// Input true to have it check, false to have it just check the block normally.
// Returns true if it is valid, false if not valid.
func (self *Wallet) VerifyBlock(block *block.Block, checkSoftwareVersion bool) bool {

	// If it is the genisis block
	if len(self.chain.Blocks) == 1 {

		return true
	}

	// Checks if the software version, if the func is told to do so
	if checkSoftwareVersion {

		if block.SoftwareVersion != util.SoftwareVersion {

			return false
		}
	}

	bytesUtil := new(util.ByteUtil)

	// Check the Block hash
	softwareVersion := []byte(block.SoftwareVersion)
	prevBlockHash, _ := hex.DecodeString(block.PrevHash)
	merkleRoot, _ := hex.DecodeString(block.MerkleRoot)
	blockTime := bytesUtil.Uint64toB(block.Timestamp)
	packedTargetBytes := bytesUtil.Uint32toB(block.PackedTarget)
	nonceBytes := bytesUtil.Uint32toB(block.Nonce)

	// Shove them together (into softwareVerion bc it is first declared)
	softwareVersion = append(softwareVersion, prevBlockHash...)
	softwareVersion = append(softwareVersion, merkleRoot...)
	softwareVersion = append(softwareVersion, packedTargetBytes...)
	softwareVersion = append(softwareVersion, blockTime...)
	softwareVersion = append(softwareVersion, nonceBytes...)

	hash := make([]byte, 32)

	// Hash the data
	sha3.ShakeSum256(hash, softwareVersion)

	// If the blockhash is invalid
	if hex.EncodeToString(hash) != block.BlockHash {

		fmt.Println(hex.EncodeToString(hash))
		return false
	}

	// Check if the block points to the previous block
	if block.PrevHash != self.chain.Blocks[self.chain.GetHeight()].BlockHash {

		return false
	}

	timeUtil := new(util.Time)

	// Check if the timestamp is valid
	// TODO: make more advanced
	if block.Timestamp < self.chain.Blocks[self.chain.GetHeight()].Timestamp || block.Timestamp > timeUtil.CurrentUnix() {

		return false
	}

	// Check if the target is correct
	if block.PackedTarget != self.chain.CalculatePackedTarget(uint(len(self.chain.Blocks))) {

		return false
	}

	// Check the merkle root
	if block.MerkleRoot != block.GetMerkleRoot() {

		return false
	}

	// Check the txs
	for index := 0; index < len(block.Txs); index += 1 {

		// If the tx is not valid
		if !self.VerifyTx(block.Txs[index]) {

			return false
		}
	}

	return true
}

// Verifys whether the blockchain attached to the wallet is valid or not.
// Returns true if valid, false if invalid.
func (self *Wallet) VerifyBlockchain() bool {

	// Is only valid if no blocks are in the chain
	if len(self.chain.Blocks) == 0 {

		return true
	}

	//****
	// Check the genisis block:

	if len(self.chain.Blocks[0].Txs) != 0 {

		return false
	}

	if self.chain.Blocks[0].PackedTarget != 0x1d0fffff {

		return false
	}

	// Check the genisis block
	//****

	//****
	// Checks the rest of the blocks

	for blockIndex := 1; blockIndex < len(self.chain.Blocks); blockIndex += 1 {

		if !self.VerifyBlock(&self.chain.Blocks[blockIndex], false) {

			return false
		}
	}

	// Checks the rest of the blocks
	//****

	return true
}
