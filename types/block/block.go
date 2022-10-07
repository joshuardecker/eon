package block

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
	"github.com/Sucks-To-Suck/Eon/types/transaction"
)

// Heads define the parameters of a block: the parents of the block, Coinbase, merkleroot of the data transactions, difficulty of the block,
// gas used (size of the data transactions), max size, creation time,
type Head struct {
	ParentHash eocrypt.Hash `json:"ParentHash"`
	Coinbase   []byte       `json:"Coinbase"`
	MerkleRoot eocrypt.Hash `json:"Merkle"`
	Difficulty *big.Int     `json:"Diff"`
	Gas        gas.Gas      `json:"GasUsed"`
	MaxGas     gas.Gas      `json:"MaxGas"`
	Time       time.Time    `json:"Time"`
	Nonce      uint         `json:"Nonce"`
	Signature  []byte       `json:"Signature"`
}

// Blocks store the data transactions of the network. Includes the header of the block, any uncles, the data transactions, an identifying hash,
// time received, and the size of the block.
type Block struct {
	Head         *Head                      `json:"Head"`
	Uncles       *[]Head                    `json:"Uncles"`
	Transactions *[]transaction.Transaction `json:"Txs"`

	BlockHash eocrypt.Hash `json:"Hash"`
	Received  time.Time    `json:"Received"`
	Size      int          `json:"Size"`
}

// ****
// Head:

// Creates and gives a Head with the given inputs.
func NewHead(ParentHash eocrypt.Hash, Coinbase []byte, Merkle eocrypt.Hash, Diff *big.Int,
	Gas gas.Gas, MaxGas gas.Gas, Time time.Time, Nonce uint) *Head {

	h := new(Head)

	h.SetParentHash(ParentHash)
	h.SetCoinbase(Coinbase)
	h.SetMerkle(Merkle)
	h.SetDiff(*Diff)
	h.SetGas(Gas)
	h.SetMaxGas(MaxGas)
	h.SetTime(Time)
	h.SetNonce(Nonce)

	return h
}

// Returns the hash of the Header.
func (h *Head) Hash() *eocrypt.Hash {

	return eocrypt.HashInterface(
		[]interface{}{

			h.ParentHash.GetBytes(),
			h.Coinbase,
			h.MerkleRoot.GetBytes(),
			h.Difficulty.Bytes(),
			h.Gas.Uint(),
			h.MaxGas.Uint(),
			h.Nonce,
			h.Signature,
		},
	)
}

func (h *Head) SetParentHash(hash eocrypt.Hash) {

	h.ParentHash = hash
}

func (h *Head) SetCoinbase(p []byte) {

	h.Coinbase = p
}

func (h *Head) SetMerkle(hash eocrypt.Hash) {

	h.MerkleRoot = hash
}

// Note about this function: none of these funcs use pointer inputs, that way data in them ant be changed unless specifically
// set by a function for safety reasons, but making a copy of the Big Int may be costly in very high throughput systems.
func (h *Head) SetDiff(b big.Int) {

	h.Difficulty = &b
}

func (h *Head) SetGas(g gas.Gas) {

	h.Gas = g
}

func (h *Head) SetMaxGas(g gas.Gas) {

	h.MaxGas = g
}

func (h *Head) SetTime(t time.Time) {

	h.Time = t
}

func (h *Head) SetNonce(n uint) {

	h.Nonce = n
}

func (h *Head) SetSig(s []byte) {

	h.Signature = s
}

func (h *Head) GetParentHashes() eocrypt.Hash { return h.ParentHash }
func (h *Head) GetCoinbase() []byte           { return h.Coinbase }
func (h *Head) GetMerkle() eocrypt.Hash       { return h.MerkleRoot }
func (h *Head) GetDiff() *big.Int             { return h.Difficulty }
func (h *Head) GetGas() gas.Gas               { return h.Gas }
func (h *Head) GetMaxGas() gas.Gas            { return h.MaxGas }
func (h *Head) GetTime() time.Time            { return h.Time }
func (h *Head) GetNonce() uint                { return h.Nonce }
func (h *Head) GetSig() []byte                { return h.Signature }

func (h *Head) Print() {

	fmt.Printf(`
	[
		Parent Hash: %x
		Coinbase: %x
		Merkle: %x
		Difficulty: %d
		Gas: %d
		Max Gas: %d	
		Time: %v
		Nonce: %v
	]
	`, h.ParentHash, h.Coinbase, h.MerkleRoot, h.Difficulty.Bytes(), h.Gas, h.MaxGas, h.Time, h.Nonce)
}

// Head:
// ****

// ****
// Block:

// Create a new block with the given inputs. This does not calculate the block hash. Call CalcHash() on the block to get the hash.
func NewBlock(Head *Head, Uncles *[]Head, Transactions *[]transaction.Transaction, ReceivedTime time.Time) *Block {

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

func (b *Block) SetHead(h Head) {

	b.Head = &h
}

func (b *Block) SetUncles(u []Head) {

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

func (b *Block) GetHead() *Head                             { return b.Head }
func (b *Block) GetUncles() *[]Head                         { return b.Uncles }
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
	`, b.CalcHash(), b.Head.ParentHash, b.Head.Coinbase, b.Head.MerkleRoot, b.Head.Difficulty.Bytes(),
		b.Head.Gas, b.Head.MaxGas, b.Head.Time, b.Head.Nonce)
}

// Block:
// ****
