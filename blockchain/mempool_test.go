package blockchain

import (
	"fmt"
	"testing"

	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

func TestMempoolCheck(t *testing.T) {

	tx1 := new(transactions.LuTx) // Valid
	tx2 := new(transactions.LuTx) // Invalid
	tx3 := new(transactions.LuTx) // Invalid, but trickier
	tx4 := new(transactions.LuTx) // Valid, but trickier

	tx1.AddScriptStr("TXID 123 OUTINDEX 456 BLKN 123", true)
	tx1.AddScriptStr("PUBK 123 SIG 456 PUBKH 789 AMT 50", false)

	tx2.AddScriptStr("TXID 123", true)
	tx2.AddScriptStr("PUBK 123", false)

	tx3.AddScriptStr("TXID 123 OUTINDEX 456", true)
	tx3.AddScriptStr("PUBK 123", false)

	tx4.AddScriptStr("OUTINDEX 456 TXID 123 COOL COOL BLKN 789", true)
	tx4.AddScriptStr("SIG 456 PUBK 123 TEST PUBKH 789 AMT 50", false)

	m := new(Mempool)

	m.AddTx(*tx1)
	m.AddTx(*tx2)
	m.AddTx(*tx3)
	m.AddTx(*tx4)

	fmt.Println("tx1:", m.CheckTxFlags(0)) // Should print true
	fmt.Println("tx2:", m.CheckTxFlags(1)) // Should print false
	fmt.Println("tx3:", m.CheckTxFlags(2)) // Should print false
	fmt.Println("tx4:", m.CheckTxFlags(3)) // Should print true
}

func TestGetTx(t *testing.T) {

	tx1 := new(transactions.LuTx) // Valid
	tx2 := new(transactions.LuTx) // Invalid
	tx3 := new(transactions.LuTx) // Valid

	tx1.AddScriptStr("BLKN 123 TXID 123 OUTINDEX 456", true)
	tx1.AddScriptStr("PUBK 123 SIG 456 PUBKH 789 AMT 50", false)

	tx2.AddScriptStr("TXID 123 TEST", true)
	tx2.AddScriptStr("PUBK 123", false)

	tx3.AddScriptStr("OUTINDEX 456 TXID 123 BLKN 456", true)
	tx3.AddScriptStr("AMT 50 PUBK 123 SIG 456 PUBKH 789", false)

	m := new(Mempool)

	m.AddTx(*tx1)
	m.AddTx(*tx2)
	m.AddTx(*tx3)

	_, err1 := m.GetTx() // Should return true
	_, err2 := m.GetTx() // Should return true
	_, err3 := m.GetTx() // Should return false

	if err1 {

		fmt.Println("Worked!")
	} else {

		fmt.Println("Didnt work!")
	}

	if err2 {

		fmt.Println("Worked!")
	} else {

		fmt.Println("Didnt work!")
	}

	if !err3 {

		fmt.Println("Worked!")
	} else {

		fmt.Println("Didnt work!")
	}
}
