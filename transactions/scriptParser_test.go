package transactions

import (
	"fmt"
	"testing"
)

func TestScripting(t *testing.T) {

	fmt.Println(ScriptParse("TXID 123 HELLO SELF HELLO HELLO"))

	for i := 0; i < 25000000; i++ {
		ScriptParse("TXID 123 HELLO SELF HELLO HELLO")
	}
}
