package merkle

import (
	"github.com/Sucks-To-Suck/Eon/eocrypt"
)

// The merkle tree is a data structure, starting off with leaf nodes of any amount greater than 0.
// From there, the FindRoot() func can be called, to mix the nodes together in pairs, and putting
// that child node into the next layer. This means each layer is half the size of the previous.
// Continues to do this until 1 node remains, which is called the root node / merkle root.
type MerkleTree struct {
	nodes map[int][]eocrypt.Hash
}

// Creates a new merkle tree with the given leaf values.
func NewMerkleTree(leafValues []eocrypt.Hash) *MerkleTree {

	// If the length of the leafs is odd, add the last leaf on the end to make it even.
	if len(leafValues)%2 == 1 {

		leafValues = append(leafValues, leafValues[len(leafValues)-1])
	}

	// Create the Merkle Tree.
	m := new(MerkleTree)

	// Initializes the map, if not done causes nil errors.
	m.nodes = make(map[int][]eocrypt.Hash)

	// Set the leaf layer (layer 0) to the given values.
	m.nodes[0] = leafValues

	return m
}

// Finds the merkle root of the Merkle Tree. Only needs the base level leafs to do this.
func (m *MerkleTree) FindRoot() *eocrypt.Hash {

	// The length of the current level, starts off with the length of the base level.
	l := len(m.nodes[0])

	// Start at base level (level 0).
	level := 0

	// Loop through until only 1 hash remains from the mixing process (the merkle root).
	for l > 1 {

		// Loop through all of the nodes in pairs, mixing those pairs into one hash thats added to the next level up.
		for nodeIndex := 0; nodeIndex < l; nodeIndex += 2 {

			m.nodes[level+1] = append(m.nodes[level+1], *mixHashes(&m.nodes[level][nodeIndex], (&m.nodes[level][nodeIndex+1])))
		}

		// Update the variables accordingly.
		l /= 2
		level += 1
	}

	// Return the merkle root.
	return &m.nodes[level][0]
}

// Mixes two given hashes together, in a very lightweight process.
func mixHashes(h1 *eocrypt.Hash, h2 *eocrypt.Hash) *eocrypt.Hash {

	// Make a byte array the size of the hash input.
	hBytes := make([]byte, len(h1.GetBytes()))

	// Loop through all of the bytes and mix (XOR) them together.
	for i, _ := range h1 {

		hBytes[i] = h1[i] ^ h2[i]
	}

	// Return it as a hash.
	return eocrypt.HashBytes(hBytes)
}
