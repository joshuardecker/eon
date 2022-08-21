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
	byteIndex := uint32(len(targetBytes) - 1)
	for targetBytes[byteIndex] == byte(0) && byteIndex > 2 {

		byteIndex -= 1
	}

	// Calculates the exponent based on how far over the byteIndex had to search
	offset := uint32(32 - len(targetBytes))
	exponent := 32 - ((byteIndex + 1) + offset)

	// Defines the packed target bytes
	pTargetBytes := make([]byte, 4)

	// Sets the value of the packed target
	pTargetBytes[0] = byte(exponent)
	pTargetBytes[1] = targetBytes[byteIndex-2]
	pTargetBytes[2] = targetBytes[byteIndex-1]
	pTargetBytes[3] = targetBytes[byteIndex]

	// Converts it to a uint32
	return binary.BigEndian.Uint32(pTargetBytes)
}

// Input the byte array of the target and this func will collapse it down into a 32bit compacted form.
// Returns the 32 bit packed form of the byte array target.
func PackTargetBytes(targetBytes []byte) uint32 {

	// Ends at 28 because this can loop 31 total times, and we do not want to pull out the significant figures (which is 3 of the 4 bytes)
	// So 31 total spots - 3 significant figures = 28, so we loop 28 times at the most.
	// It will also stop when it hits a non 0 value in the target, aka when it hits a significant figure.
	byteIndex := uint32(len(targetBytes) - 1)
	for targetBytes[byteIndex] == byte(0) && byteIndex > 2 {

		byteIndex -= 1
	}

	// Calculates the exponent based on how far over the byteIndex had to search
	offset := uint32(32 - len(targetBytes))
	exponent := 32 - ((byteIndex + 1) + offset)

	// Defines the packed target bytes
	pTargetBytes := make([]byte, 4)

	// Sets the value of the packed target
	pTargetBytes[0] = byte(exponent)
	pTargetBytes[1] = targetBytes[byteIndex-2]
	pTargetBytes[2] = targetBytes[byteIndex-1]
	pTargetBytes[3] = targetBytes[byteIndex]

	// Converts it to a uint32
	return binary.BigEndian.Uint32(pTargetBytes)
}
