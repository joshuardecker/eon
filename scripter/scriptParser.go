package scripter

import (
	"strings"
)

// Script flags that will have a proceeding value
var keyWordsValuePair []string = []string{}

// Script flags without a proceeding value
var keyWordsNoPair []string = []string{}

// This function takes in a script string and parses it.
// It will remove junk that isnt a valid flag or value after a flag.
// It also sorts each script flag and value into an individual spot in an array.
// Example if inputted "TXID 123", TXID will be in index 0 of the array, and 123 at index 1.
// Returns the string array of the parsed script with junk removed.
func StrToScript(inputScript string) []string {

	// This is the input script thats had everything moved into an array. Sorts each item by a " " occuring.
	// Example "test1 test2". Test1 is now in the array at index 0, and test2 is at index 1 because it was split at the " ".
	args := strings.Split(inputScript, " ")

	// Save the length so it doesnt have to be calculated every loop of the for loop below.
	leng := uint(len(args))

	// Init the index
	index := uint(0)

	for index < leng {

		switch checkKeyWord(args[index]) {

		// If the string is not a flag
		case 0:
			args = removeString(args, index)

			leng -= 1

		// If the string is a flag that needs a proceeding value
		case 1:

			if leng-index == 1 {
				// If there is only one value left, aka the flag has no coorisponding value

				args = removeString(args, index)
				leng -= 1

			} else if checkKeyWord(args[index+1]) == 1 {
				// If the value of this flag is another flag, aka invalid

				args = removeString(args, index)
				leng -= 1
			} else {
				// If everything is good, include this flag and its value

				index += 2
			}

		// If the string is a flag with no proceeding value
		case 2:
			index += 1
		}
	}

	args = removeCopyFlags(args)

	return args
}

// Removes a string at the given index.
// Returns this new string array.
func removeString(inputArgs []string, index uint) []string {

	return append(inputArgs[:index], inputArgs[index+1:]...)
}

// This function is given a string and determins whether it is a valid flag or not.
// Returns 0 if the flag is not valid.
// Returns 1 if the flag is valid and has a proceeding value.
// Returns 2 if the flag is valid and has no proceeding value.
func checkKeyWord(word string) uint {

	// Saves the length so it doesnt have to be recalculated every loop
	leng := len(keyWordsValuePair)

	// Checks through the list of flags with proceeding values
	for index := 0; index < leng; index++ {

		// It is a flag, and has a proceeding value
		if word == keyWordsValuePair[index] {

			return 1
		}
	}

	// Saves the length so it doesnt have to be recalculated every loop
	leng = len(keyWordsNoPair)

	// Checks through the list of flags with no proceeding values
	for index := 0; index < leng; index++ {

		// It is a flag, but has no proceeding value
		if word == keyWordsNoPair[index] {

			return 2
		}
	}

	return 0
}

// Removes all of the copy flags.
// Starts from index 0, so the flags closest to index 0 will be kept,
// the extra copys removed from the script.
// Input is the script.
// Output is the script without the extra of the same flags.
func removeCopyFlags(script []string) []string {

	// Used to identify if the same flag has been used
	var flagFound bool

	// Saves the length so it doesnt need to be recalculated
	leng := len(keyWordsValuePair)
	scriptLeng := len(script)

	// Check all of the words in keyWordsValuePair
	for index := 0; index < leng; index += 1 {

		// Checks the whole script
		for nIndex := 0; nIndex < scriptLeng; nIndex += 1 {

			if keyWordsValuePair[index] == script[nIndex] && !flagFound {
				// When the flag is used the first time

				flagFound = true
			} else if keyWordsValuePair[index] == script[nIndex] && flagFound {
				// When the same flag is used again

				// Remove the copy and its value pair
				script = removeString(script, uint(nIndex))
				script = removeString(script, uint(nIndex))

				scriptLeng -= 2

				nIndex -= 1
			}
		}

		flagFound = false
	}

	// Saves the length so it doesnt need to be recalculated
	leng = len(keyWordsNoPair)

	// Check all of the words in keyWordsValuePair
	for index := 0; index < leng; index += 1 {

		// Checks the whole script
		for nIndex := 0; nIndex < scriptLeng; nIndex += 1 {

			if keyWordsNoPair[index] == script[nIndex] && !flagFound {
				// When the flag is used the first time

				flagFound = true
			} else if keyWordsNoPair[index] == script[nIndex] && flagFound {
				// When the same flag is used again

				// Remove the extra flag
				script = removeString(script, uint(nIndex))

				scriptLeng -= 1

				nIndex -= 1
			}
		}

		flagFound = false
	}

	return script
}

// Function takes a given script and converts it back into a single string.
// Input is a script.
// Returns this string.
func ScriptToStr(script []string) string {

	// Saves the length so it doesnt have to be recalculated every loop
	leng := len(script)

	// Define the script string
	var scriptStr string

	// Comvert the script into the script string
	for index := 0; index < leng; index++ {

		scriptStr += (script[index] + " ")
	}

	return scriptStr
}
