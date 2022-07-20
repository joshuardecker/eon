package transactions

import "strings"

// Key words:
// For input: TXID (the tx id of where you get the coin), OUTINDEX (which index in that tx), PUBKH (the receivers (you) hashed PubKey)
// For output: PUBK (your public key), SIG (your signature), PUBKH (the hash of the public key you are sending coin to), AMT (amount of Luncheon you are sending)
// Extras:
// SELF (sends the remaining inputted coin to yourself), TIML (time lock, how long the receiver must wait to spend there coin)

// Flags that will have a value after it
var keyWordsValuePair []string = []string{
	"TXID",
	"OUTINDEX",
	"PUBKH",
	"PUBK",
	"SIG",
	"AMT",
	"TIML",
}

// Flags without a value after it
var keyWordsNoPair []string = []string{
	"SELF",
}

func ScriptParse(inputScript string) []string {

	args := strings.Split(inputScript, " ")
	leng := uint(len(args))
	index := uint(0)

	for index < leng {

		switch checkKeyWord(args[index]) {

		case 0:
			args = removeString(args, index)

			leng -= 1
		case 1:
			index += 2

		case 2:
			index += 1
		}
	}

	return args
}

func removeString(inputArgs []string, index uint) []string {

	return append(inputArgs[:index], inputArgs[index+1:]...)
}

// The first return is if the input is a key word or not.
// The second return is if the proceeding string is a value.
// True if it is, false if it doesnt need a proceeding value
func checkKeyWord(word string) uint {

	leng := len(keyWordsValuePair)
	for index := 0; index < leng; index++ {

		// It is a flag, and has a proceeding value
		if word == keyWordsValuePair[index] {

			return 1
		}
	}

	leng = len(keyWordsNoPair)
	for index := 0; index < leng; index++ {

		if word == keyWordsNoPair[index] {

			return 2
		}
	}

	return 0
}
