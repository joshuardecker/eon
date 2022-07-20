package transactions

import (
	"fmt"
	"testing"
)

func TestScripting(t *testing.T) {

	str := "TXID 123 HELLO SELF HELLO HELLO TXID 123 TXID 123 SELF HELLO"

	script := ScriptParse(str)
	fmt.Println("Script: ", script)

	scriptStr := ScriptToStr(script)
	fmt.Println("Script Str: ", scriptStr)

	script = ScriptParse(scriptStr)
	fmt.Println("Script: ", script)

	scriptStr = ScriptToStr(script)
	fmt.Println("Script Str: ", scriptStr)
}
