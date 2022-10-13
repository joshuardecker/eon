package transaction

import (
	"math/big"
	"testing"

	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
	"github.com/Sucks-To-Suck/Eon/tools/eotime"
)

func TestTransaction(t *testing.T) {

	// Create the variables for the tx:
	tokenHash := eocrypt.HashBytes([]byte("Lunch Coin!"))
	amount := big.NewInt(1)
	to := []byte("kaimort123")
	from := []byte("the-ol-jd")
	sig := []byte("Totallt Legit Sig")
	blockFrom := make([]eocrypt.Hash, 1)
	blockFrom = append(blockFrom, *eocrypt.HashBytes([]byte("100")))
	txFrom := make([]eocrypt.Hash, 1)
	txFrom = append(txFrom, *eocrypt.HashBytes([]byte("my-cool-tx")))
	chain := big.NewInt(100)
	txGas := gas.Gas(1)
	gasPrice := big.NewInt(1)
	txTime := eotime.LocalTime()

	// Create the transaction.
	tx := NewTransaction(*tokenHash, amount, to, from, sig, blockFrom,
		txFrom, chain, txGas, *gasPrice, txTime)

	// Print the transaction.
	tx.Print()
}
