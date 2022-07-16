package luncheon

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
