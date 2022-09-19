package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/gob"
	"encoding/json"
	"math/big"
	"time"

	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
)

// A Transaction simply represents a token being sent from one party to another (not even the native token).
type Transaction struct {
	// Token data of the transaction.
	TokenHash eocrypt.Hash `json:"TokenHash"`
	Amount    *big.Int     `json:"Amount"`

	// From who and where to?
	To        *ecdsa.PublicKey `json:"To"`
	From      *ecdsa.PublicKey `json:"From"`
	Signature []byte           `json:"Signature"`

	// Where are you getting the tokens from?
	BlockFrom []eocrypt.Hash `json:"BlockFromHash"`
	TxFrom    []eocrypt.Hash `json:"TransactionFromHash"`

	// Basic but otherwise needed data.
	ChainId  *big.Int     `json:"ChainId"`
	Gas      gas.Gas      `json:"Gas"`
	GasPrice gas.GasPrice `json:"GasPrice"`

	// When did the node get this transaction?
	ReceivedTime time.Time `json:"ReceivedTime"`

	// Calculated data about the transaction.
	Hash eocrypt.Hash `json:"Hash"`
	Size int          `json:"Size"`
}

// ****
// Transaction basic interaction:

// The following functions all dont use / give pointers.
// This is to allow manipulation of the given / received data that then will not affect the copy in the tx until a function
// is called to set the new value.

func NewTransaction(TokenHash eocrypt.Hash, Amount *big.Int, To *ecdsa.PublicKey, From *ecdsa.PublicKey, Signature []byte,
	BlockFrom []eocrypt.Hash, TxFrom []eocrypt.Hash, ChainId *big.Int, Gas gas.Gas, GasPrice gas.GasPrice,
	ReceivedTime time.Time) *Transaction {

	t := new(Transaction)

	t.SetTokenHash(TokenHash)
	t.SetAmount(*Amount)
	t.SetTo(*To)
	t.SetFrom(*From)
	t.SetSignature(Signature)
	t.SetBlockFrom(BlockFrom)
	t.SetTxFrom(TxFrom)
	t.SetChainId(*ChainId)
	t.SetGas(Gas)
	t.SetGasPrice(GasPrice)
	t.SetReceivedTime(ReceivedTime)

	return t
}

func (t *Transaction) SetTokenHash(h eocrypt.Hash) {

	t.TokenHash = h
}

func (t *Transaction) SetAmount(b big.Int) {

	t.Amount = &b
}

func (t *Transaction) SetTo(p ecdsa.PublicKey) {

	t.To = &p
}

func (t *Transaction) SetFrom(p ecdsa.PublicKey) {

	t.From = &p
}

func (t *Transaction) SetBlockFrom(h []eocrypt.Hash) {

	t.BlockFrom = h
}

func (t *Transaction) SetTxFrom(h []eocrypt.Hash) {

	t.TxFrom = h
}

func (t *Transaction) SetSignature(sig []byte) {

	t.Signature = sig
}

func (t *Transaction) SetChainId(b big.Int) {

	t.ChainId = &b
}

func (t *Transaction) SetGas(g gas.Gas) {

	t.Gas = g
}

func (t *Transaction) SetGasPrice(g gas.GasPrice) {

	t.GasPrice = g
}

func (t *Transaction) SetReceivedTime(time time.Time) {

	t.ReceivedTime = time
}

func (t *Transaction) SetHash(h eocrypt.Hash) {

	t.Hash = h
}

func (t *Transaction) SetSize(s int) {

	t.Size = s
}

func (t *Transaction) GetTokenHash() eocrypt.Hash   { return t.TokenHash }
func (t *Transaction) GetAmount() big.Int           { return *t.Amount }
func (t *Transaction) GetTo() ecdsa.PublicKey       { return *t.To }
func (t *Transaction) GetFrom() ecdsa.PublicKey     { return *t.From }
func (t *Transaction) GetBlockFrom() []eocrypt.Hash { return t.BlockFrom }
func (t *Transaction) GetTxFrom() []eocrypt.Hash    { return t.TxFrom }
func (t *Transaction) GetChainId() big.Int          { return *t.ChainId }
func (t *Transaction) GetGas() gas.Gas              { return t.Gas }
func (t *Transaction) GetGasPrice() gas.GasPrice    { return t.GasPrice }
func (t *Transaction) GetReceivedTime() time.Time   { return t.ReceivedTime }
func (t *Transaction) GetHash() eocrypt.Hash        { return t.Hash }
func (t *Transaction) GetSize() int                 { return t.Size }

// Transaction basic interaction:
// ****

// ****
// Transaction advanced interaction:

// Uses the gob encoder to encode and return the transaction as a Bytes Buffer.
func (t *Transaction) EncodeToBuffer() (*bytes.Buffer, error) {

	buff := new(bytes.Buffer)

	// Encode the transaction into the Bytes Buffer.
	encodeErr := gob.NewEncoder(buff).Encode(t)

	return buff, encodeErr
}

// Uses the gob encoder with a provided Bytes Buffer to encode the transaction into that Buffer.
func (t *Transaction) EncodeWithBuffer(b *bytes.Buffer) error {

	// Encode the transaction into the Bytes Buffer.
	return gob.NewEncoder(b).Encode(t)
}

// Encode the transaction into JSON format. Returns the encoded bytes in a Bytes Buffer.
func (t *Transaction) EncodeJSON() (*bytes.Buffer, error) {

	buff := new(bytes.Buffer)

	// Encode the transaction into the Bytes Buffer.
	encodeErr := json.NewEncoder(buff).Encode(t)

	return buff, encodeErr
}

// Encode the transaction with the provided Bytes Buffer. The encoded bytes will reside there.
func (t *Transaction) EncodeJSONwithBuff(b *bytes.Buffer) error {

	// Encode the transaction into the Bytes Buffer.
	return json.NewEncoder(b).Encode(t)
}

// Trys to decode the given Bytes Buffer (suppose to be the encoded form by gob). Returns the decoded transaction and any errors.
func Decode(b *bytes.Buffer) (*Transaction, error) {

	t := new(Transaction)

	// Try to decode the Bytes Buffer into a transaction.
	decodeErr := gob.NewDecoder(b).Decode(t)

	return t, decodeErr
}

// Trys to decode the given Bytes Buffer (suppose to be encoded JSON form). Returns the decoded transaction and any errors.
func DecodeJSON(b *bytes.Buffer) (*Transaction, error) {

	t := new(Transaction)

	// Try to decode the Bytes Buffer into a transaction.
	decodeErr := json.NewDecoder(b).Decode(t)

	return t, decodeErr
}

// Transaction advanced interaction:
// ****
