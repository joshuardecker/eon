package target

import (
	"encoding/binary"
	"math/big"
)

// Input the bigInt target and this func will collapse it down into a 32bit compacted form.
// Returns that compacted form.
func PackTargetInt(target *big.Int) uint32 {

	targetBytes := target.Bytes()

	// Ends at 28 because this can loop 31 total times, and we do not want to pull out the significant figures (which is 3 of the 4 bytes)
	// So 31 total spots - 3 significant figures = 28, so we loop 28 times at the most.
	// It will also stop when it hits a non 0 value in the target, aka when it hits a significant figure.
	byteIndex := uint32(0)
	for byteIndex = 0; targetBytes[byteIndex] == byte(0) && byteIndex < 28; byteIndex += 1 {
	}

	// Calculates the exponent based on how far over the byteIndex had to search

	// Offset is determined by how many 0's were left out from the most left of the target.
	// Example, if 003654 was saved as 3654, the offset would be 2, because 2 0's were not saved at the front of the number.
	offset := uint32(32 - len(targetBytes))

	// Exponent is calculated as the max exponent (32 as 32 bytes) - how many times the byteIndex moved to hit a sig fig, plus the offset calculated above.
	// Finally plus 3 because we need to include the 3 sig figs that we are saving, as we look for the left most one, not the right most sig fig.
	exponent := 32 - (byteIndex + offset + 3)

	// Defines the packed target bytes
	pTargetBytes := make([]byte, 4)

	// Sets the value of the packed target
	pTargetBytes[0] = byte(exponent)
	pTargetBytes[1] = targetBytes[byteIndex]
	pTargetBytes[2] = targetBytes[byteIndex+1]
	pTargetBytes[3] = targetBytes[byteIndex+2]

	// Converts it to a uint32
	return binary.BigEndian.Uint32(pTargetBytes)
}

// Input the byte array of the target and this func will collapse it down into a 32bit compacted form.
// Returns the 32 bit packed form of the byte array target.
func PackTargetBytes(targetBytes []byte) uint32 {

	// Ends at 28 because this can loop 31 total times, and we do not want to pull out the significant figures (which is 3 of the 4 bytes)
	// So 31 total spots - 3 significant figures = 28, so we loop 28 times at the most.
	// It will also stop when it hits a non 0 value in the target, aka when it hits a significant figure.
	byteIndex := uint32(0)
	for byteIndex = 0; targetBytes[byteIndex] == byte(0) && byteIndex < 28; byteIndex += 1 {
	}

	// Calculates the exponent based on how far over the byteIndex had to search

	// Offset is determined by how many 0's were left out from the most left of the target.
	// Example, if 003654 was saved as 3654, the offset would be 2, because 2 0's were not saved at the front of the number.
	offset := uint32(32 - len(targetBytes))

	// Exponent is calculated as the max exponent (32 as 32 bytes) - how many times the byteIndex moved to hit a sig fig, plus the offset calculated above.
	// Finally plus 3 because we need to include the 3 sig figs that we are saving, as we look for the left most one, not the right most sig fig.
	exponent := 32 - (byteIndex + offset + 3)

	// Defines the packed target bytes
	pTargetBytes := make([]byte, 4)

	// Sets the value of the packed target
	pTargetBytes[0] = byte(exponent)
	pTargetBytes[1] = targetBytes[byteIndex]
	pTargetBytes[2] = targetBytes[byteIndex+1]
	pTargetBytes[3] = targetBytes[byteIndex+2]

	// Converts it to a uint32
	return binary.BigEndian.Uint32(pTargetBytes)
}
