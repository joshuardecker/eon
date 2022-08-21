package block

import (
	"fmt"
	"testing"

	"github.com/Sucks-To-Suck/LuncheonNetwork/core/events/txs"
	"golang.org/x/crypto/sha3"
)

func TestMerkle(t *testing.T) {

	block := new(Block)

	tx1 := new(txs.LuTx)
	tx2 := new(txs.LuTx)
	tx3 := new(txs.LuTx)

	tx1.AddScriptStr("PUBKH 123")
	tx2.AddScriptStr("PUBKH 124")
	tx3.AddScriptStr("PUBKH 125")

	// By hand way

	tx1Bytes := tx1.AsBytes()
	tx2Bytes := tx2.AsBytes()
	tx3Bytes := tx3.AsBytes()

	tx1Bytes = append(tx1Bytes, tx2Bytes...)
	tx3Bytes = append(tx3Bytes, tx3Bytes...)

	hash1 := make([]byte, 32)
	hash2 := make([]byte, 32)

	sha3.ShakeSum256(hash1, tx1Bytes)
	sha3.ShakeSum256(hash2, tx3Bytes)

	hash1 = append(hash1, hash2...)

	hash3 := make([]byte, 32)
	sha3.ShakeSum256(hash3, hash1)

	fmt.Printf("Merkle Root: %x\n", hash3)

	// Actual Way

	block.Txs = append(block.Txs, *tx1)
	block.Txs = append(block.Txs, *tx2)
	block.Txs = append(block.Txs, *tx3)

	fmt.Println("Merkle Root:", block.GetMerkleRoot())
}
