package transactions

import (
	"fmt"
	"testing"
)

func TestScripting(t *testing.T) {

	str := "TXID 123 HELLO SELF HELLO HELLO TXID 123 TXID 123 SELF HELLO"

	script := StrToScript(str)
	fmt.Println("Script: ", script)

	scriptStr := ScriptToStr(script)
	fmt.Println("Script Str: ", scriptStr)

	script = StrToScript(scriptStr)
	fmt.Println("Script: ", script)

	scriptStr = ScriptToStr(script)
	fmt.Println("Script Str: ", scriptStr)
}

func TestScriptAdvanced(t *testing.T) {

	script1 := StrToScript("TXID 123 AMT 123 SELF")  // Valid script
	script2 := StrToScript("TXID 123 AMT 123 HELLO") // Will need HELLO removed to be valid
	script3 := StrToScript("TXID AMT 123 SELF")      // Needs TXID removed to be valid

	// Print the valid scripts
	fmt.Println("1:", ScriptToStr(script1))
	fmt.Println("2:", ScriptToStr(script2))
	fmt.Println("3:", ScriptToStr(script3))
}
