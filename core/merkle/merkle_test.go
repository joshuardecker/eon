package merkle

import (
	"fmt"
	"testing"

	"github.com/Sucks-To-Suck/Eon/eocrypt"
)

func TestMerkle(t *testing.T) {

	// Test if it works:
	leaf := make([]eocrypt.Hash, 4)

	leaf[0] = *eocrypt.HashBytes([]byte("Hello World."))
	leaf[1] = *eocrypt.HashBytes([]byte("Hello World!"))
	leaf[2] = *eocrypt.HashBytes([]byte("Hello World!!"))
	leaf[3] = *eocrypt.HashBytes([]byte("Hello World?"))

	merkle := NewMerkleTree(leaf)
	fmt.Printf("%x\n", *merkle.FindRoot())

	// Test if slightly changing the input way changes the output:
	leaf[3] = *eocrypt.HashBytes([]byte("Hello World??"))

	merkle = NewMerkleTree(leaf)
	fmt.Printf("%x\n", *merkle.FindRoot())

	// Test if odd amounts of leafs works:
	newLeaf := make([]eocrypt.Hash, 3)
	newLeaf[0] = *eocrypt.HashBytes([]byte("Fun fact #1"))
	newLeaf[1] = *eocrypt.HashBytes([]byte("Fun fact #2"))
	newLeaf[2] = *eocrypt.HashBytes([]byte("Fun fact #3"))

	merkle = NewMerkleTree(newLeaf)
	fmt.Printf("%x\n", *merkle.FindRoot())
}
