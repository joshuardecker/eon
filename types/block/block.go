package block

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/gob"
	"encoding/json"
	"time"

	"github.com/Sucks-To-Suck/Eon/eocrypt"
	"github.com/Sucks-To-Suck/Eon/eocrypt/curve"
	"github.com/Sucks-To-Suck/Eon/types/transaction"
)

// A Block is a storgae of a defining Header, any uncle blocks, and the data transactions within.
type Block struct {
	header       Header                             `json:"Head"`
	uncles       []Header                           `json:"Uncles"`
	transactions map[string]transaction.Transaction `json:"Txs"`

	hash     eocrypt.Hash `json:"Hash"`
	received time.Time    `json:"Received"`
}

// Creates and returns a new block with the given inputs.
func NewBlock(Head Header, Uncles []Header, Transactions []transaction.Transaction, ReceivedTime time.Time) *Block {

	b := new(Block)

	// Apply the Block data (minus the transactions).
	b.header = Head
	b.uncles = Uncles
	b.received = ReceivedTime

	// Loop through and add all of the txs to the map, where the index is the tx hash as a string, and the value is the tx.
	for _, tx := range Transactions {

		b.transactions[tx.Hash().String()] = tx
	}

	return b
}

// Calculates and stores the Hash of the block. Also returns the hash of the block.
func (b *Block) CalcHash() eocrypt.Hash {

	b.hash = b.header.Hash()

	return b.hash
}

func (b *Block) Header() Header                                   { return b.header }
func (b *Block) Uncles() []Header                                 { return b.uncles }
func (b *Block) Transactions() map[string]transaction.Transaction { return b.transactions }
func (b *Block) Time() time.Time                                  { return b.received }

// Note this does not calculate the hash, just returns the stored value.
// Use CalcHash() to calculate and get the block hash.
func (b *Block) Hash() eocrypt.Hash { return b.hash }

// Uses the gob encoder with a provided Bytes Buffer to encode the block into that Buffer.
func (b *Block) EncodeWithBuffer(by *bytes.Buffer) error {

	// Encode the block into the Bytes Buffer.
	return gob.NewEncoder(by).Encode(b)
}

// Encode the block with the provided Bytes Buffer. The encoded bytes will reside there.
func (b *Block) EncodeJSON(by *bytes.Buffer) error {

	// Encode the block into the Bytes Buffer.
	return json.NewEncoder(by).Encode(b)
}

// Trys to decode the given Bytes Buffer (suppose to be the encoded form by gob). Returns the decoded block and any errors.
func Decode(by *bytes.Buffer) (*Block, error) {

	b := new(Block)

	// Try to decode the Bytes Buffer into a block.
	decodeErr := gob.NewDecoder(by).Decode(b)

	return b, decodeErr
}

// Trys to decode the given Bytes Buffer (suppose to be encoded JSON form). Returns the decoded block and any errors.
func DecodeJSON(by *bytes.Buffer) (*Block, error) {

	b := new(Block)

	// Try to decode the Bytes Buffer into a block.
	decodeErr := json.NewDecoder(by).Decode(b)

	return b, decodeErr
}

// Returns the encoded bytes of the block.
func (b *Block) Bytes() []byte {

	// Create the bytes buffer.
	buff := new(bytes.Buffer)

	// Encode the block into the buffer.
	b.EncodeWithBuffer(buff)

	return buff.Bytes()
}

// Prints the Header of the block.
func (b *Block) Print() {

	b.header.PrintBlock(b.hash)
}

// Signs the block with the given private key. Returns any errors.
func (b *Block) Sign(key *ecdsa.PrivateKey) error {

	// Get the signature and any errors.
	sig, err := curve.Sign(key, b.hash.Bytes())

	// Put the signature into the header.
	b.header.Sign(sig)

	// Return any errors.
	return err
}
