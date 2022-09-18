package transaction

import (
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
)

// A Transaction simply represents a token being sent from one party to another (not even the native token).
type Transaction struct {
	// Token data of the transaction.
	TokenHash eocrypt.Hash
	Amount    *big.Int

	// From who and where to?
	To        *ecdsa.PublicKey
	From      *ecdsa.PublicKey
	Signature []byte

	// Where are you getting the tokens from?
	BlockFrom eocrypt.Hash
	TxFrom    eocrypt.Hash

	// Basic but otherwise needed data.
	ChainId  *big.Int
	Gas      gas.Gas
	GasPrice gas.GasPrice

	// When did the node get this transaction?
	ReceivedTime time.Time

	// Calculated data about the transaction.
	Hash eocrypt.Hash
	Size int
}

// ****
// Transaction basic interaction:

// The following functions all dont use / give pointers.
// This is to allow manipulation of the given / received data that then will not affect the copy in the tx until a function
// is called to set the new value.

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

func (t *Transaction) GetTokenHash() eocrypt.Hash { return t.TokenHash }
func (t *Transaction) GetAmount() big.Int         { return *t.Amount }
func (t *Transaction) GetTo() ecdsa.PublicKey     { return *t.To }
func (t *Transaction) GetFrom() ecdsa.PublicKey   { return *t.From }
func (t *Transaction) GetChainId() big.Int        { return *t.ChainId }
func (t *Transaction) GetGas() gas.Gas            { return t.Gas }
func (t *Transaction) GetGasPrice() gas.GasPrice  { return t.GasPrice }
func (t *Transaction) GetReceivedTime() time.Time { return t.ReceivedTime }
func (t *Transaction) GetHash() eocrypt.Hash      { return t.Hash }
func (t *Transaction) GetSize() int               { return t.Size }

// Transaction basic interaction:
// ****

// ****
// Transaction advanced interaction:

// Transaction advanced interaction:
// ****
