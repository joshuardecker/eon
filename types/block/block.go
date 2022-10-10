package block

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Sucks-To-Suck/Eon/eocrypt"
	"github.com/Sucks-To-Suck/Eon/types/transaction"
)

// Blocks store the data transactions of the network. Includes the header of the block, any uncles, the data transactions, an identifying hash,
// time received, and the size of the block.
type Block struct {
	Head         *Header                    `json:"Head"`
	Uncles       *[]Header                  `json:"Uncles"`
	Transactions *[]transaction.Transaction `json:"Txs"`

	BlockHash eocrypt.Hash `json:"Hash"`
	Received  time.Time    `json:"Received"`
	Size      int          `json:"Size"`
}

// ****
// Block:

// Create a new block with the given inputs. This does not calculate the block hash. Call CalcHash() on the block to get the hash.
func NewBlock(Head *Header, Uncles *[]Header, Transactions *[]transaction.Transaction, ReceivedTime time.Time) *Block {

	b := new(Block)

	b.SetHead(*Head)

	// If there are no uncles, do not set anything.
	if Uncles != nil {

		b.SetUncles(*Uncles)
	}

	// If there are no transactions, do nothing.
	if Transactions != nil {

		b.SetTransactions(*Transactions)
	}

	b.SetReceivedTime(ReceivedTime)

	return b
}

// This function calculates the hash of the block and saves it in the block. It also returns the hash, allowing for a copy to be
// interacted with.
func (b *Block) CalcHash() eocrypt.Hash {

	b.SetHash(*b.Head.Hash())

	return b.GetHash()
}

func (b *Block) SetHead(h Header) {

	b.Head = &h
}

func (b *Block) SetUncles(u []Header) {

	b.Uncles = &u
}

func (b *Block) SetTransactions(t []transaction.Transaction) {

	b.Transactions = &t
}

func (b *Block) SetHash(h eocrypt.Hash) {

	b.BlockHash = h
}

func (b *Block) SetReceivedTime(t time.Time) {

	b.Received = t
}

func (b *Block) GetHead() *Header                           { return b.Head }
func (b *Block) GetUncles() *[]Header                       { return b.Uncles }
func (b *Block) GetTransactions() []transaction.Transaction { return *b.Transactions }
func (b *Block) GetHash() eocrypt.Hash                      { return b.BlockHash }
func (b *Block) GetReceivedTime() time.Time                 { return b.Received }

// Uses the gob encoder to encode and return the block as a Bytes Buffer.
func (b *Block) EncodeToBuffer() (*bytes.Buffer, error) {

	buff := new(bytes.Buffer)

	// Encode the block into the Bytes Buffer.
	encodeErr := gob.NewEncoder(buff).Encode(b)

	return buff, encodeErr
}

// Uses the gob encoder with a provided Bytes Buffer to encode the block into that Buffer.
func (b *Block) EncodeWithBuffer(by *bytes.Buffer) error {

	// Encode the block into the Bytes Buffer.
	return gob.NewEncoder(by).Encode(b)
}

// Encode the block into JSON format. Returns the encoded bytes in a Bytes Buffer.
func (b *Block) EncodeJSON() (*bytes.Buffer, error) {

	buff := new(bytes.Buffer)

	// Encode the block into the Bytes Buffer.
	encodeErr := json.NewEncoder(buff).Encode(b)

	return buff, encodeErr
}

// Encode the block with the provided Bytes Buffer. The encoded bytes will reside there.
func (b *Block) EncodeJSONwithBuff(by *bytes.Buffer) error {

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

	bytes, _ := b.EncodeToBuffer()

	return bytes.Bytes()
}

func (b *Block) Print() {
	fmt.Printf(`
	[
		Block Hash: %x
	
		Parent Hash: %x
		Coinbase: %x
		Merkle: %x
		Difficulty: %d
		Gas: %d
		Max Gas: %d	
		Time: %v
		Nonce: %v
	]
	`, b.CalcHash(), b.Head.ParentHash(), b.Head.Coinbase(), b.Head.MerkleRoot(), b.Head.Difficulty().Bytes(),
		b.Head.Gas(), b.Head.GasLimit(), b.Head.Time(), b.Head.Nonce())
}

// Block:
// ****
