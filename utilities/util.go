package utilities

// Provides basic utilities that will be needed throughout the program.
// Example converting uint32's to a byte array.
type Util struct{}

// Converts a given uint32 into a byte array.
// Returns sed byte array.
func (u Util) Uint32toB(inputUint uint32) []byte {

	// Init the byte array
	bArray := make([]byte, 4)

	// Loops to correctly bitshift and mask to copy into the byte array.
	for i := uint32(0); i < 4; i++ {

		bArray[i] = byte((inputUint >> (8 * i)) & 0xff)
	}

	return bArray
}

// Converts a given uint64 into a byte array.
// Returns sed byte array.
func (u Util) Uint64toB(inputUint uint64) []byte {

	// Init the byte array
	bArray := make([]byte, 8)

	// Loops to correctly bitshift and mask to copy into the byte array.
	for i := uint32(0); i < 8; i++ {

		bArray[i] = byte((inputUint >> (8 * i)) & 0xff)
	}

	return bArray
}

// Flips / swaps the order of the inputted byte array. Returns sed flipped byte array.
func (u Util) ByteArraySwap(byteArray []byte) []byte {

	for swapStart := 0; swapStart < len(byteArray)/2; swapStart++ {

		// Defines the correct swapEnd index based on the swapStart index value
		swapEnd := len(byteArray) - swapStart - 1 // -1 is because indexes are 1 less than the length, ex: length 8's last index is [7]

		// Swaps the values
		byteArray[swapStart], byteArray[swapEnd] = byteArray[swapEnd], byteArray[swapStart]
	}

	return byteArray
}
