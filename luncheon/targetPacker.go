package luncheon

import (
	"encoding/binary"
	"errors"

	"github.com/GoblinBear/beson/types"
)

// A struct that handles converting the 256bit target into its 32bit, packed form.
// Can handle the uint256 or the byte array of that uint as valid inputs.
type TargetPacker struct {
	unpackedTarget      types.UInt256
	unpackedTargetBytes []byte

	packedTarget uint32
	exponent     uint32
	byteIndex    uint32

	util Util
}

// Takes in a uint256 and returns its packed 32 bit form.
func (t *TargetPacker) PackTargetUint256(unpackedTarget types.UInt256) uint32 {

	// Init the packed target with a value of 0
	t.packedTarget = 0

	// Save variables into struct for later use
	t.unpackedTarget = unpackedTarget
	t.unpackedTargetBytes = t.unpackedTarget.Get()

	// Loop through the byte array until it hits a non 0 value
	for t.byteIndex = uint32(0); t.unpackedTargetBytes[t.byteIndex] == byte(0); t.byteIndex++ {

		// Prevents looping out of valid index (31 - 3 = 28) as three values will be pulled out
		if t.byteIndex == 28 {

			break
		}
	}

	// Calculates the exponent based on how far over the byteIndex had to search
	t.exponent = 32 - t.byteIndex

	// Defines the byte array
	byteArray := make([]byte, 4)

	// Sets the value of the packed target
	byteArray[0] = byte(t.exponent)
	byteArray[1] = t.unpackedTargetBytes[t.byteIndex]
	byteArray[2] = t.unpackedTargetBytes[t.byteIndex+1]
	byteArray[3] = t.unpackedTargetBytes[t.byteIndex+2]

	// Converts it to a uint32
	t.packedTarget = binary.BigEndian.Uint32(byteArray)

	return t.packedTarget
}

// Takes in a byte array of the uint256 and returns its packed 32 bit form.
// Returns errors if the byte array inputted is not big enough.
func (t *TargetPacker) PackTargetBytes(unpackedTargetBytes []byte) (uint32, error) {

	// Init the packed target with a value of 0
	t.packedTarget = 0

	// If the inputted bytes are not long enough (not a uint256)
	if len(unpackedTargetBytes) < 32 {
		return t.packedTarget, errors.New("Input byte array not long enough!")
	}

	// Save variables into struct for later use
	t.unpackedTarget.Set(unpackedTargetBytes)
	t.unpackedTargetBytes = unpackedTargetBytes

	// Loop through the byte array until it hits a non 0 value
	for t.byteIndex = 0; t.unpackedTargetBytes[t.byteIndex] == byte(0); t.byteIndex++ {

		// Prevents looping out of valid index (31 - 3 = 28) as three values will be pulled out
		if t.byteIndex == 28 {

			break
		}
	}

	// Calculates the exponent based on how far over the byteIndex had to search
	t.exponent = 32 - t.byteIndex

	// Defines the byte array
	byteArray := make([]byte, 4)

	// Sets the value of the packed target
	byteArray[0] = byte(t.exponent)
	byteArray[1] = t.unpackedTargetBytes[t.byteIndex]
	byteArray[2] = t.unpackedTargetBytes[t.byteIndex+1]
	byteArray[3] = t.unpackedTargetBytes[t.byteIndex+2]

	// Converts it to a uint32
	t.packedTarget = binary.BigEndian.Uint32(byteArray)

	return t.packedTarget, nil // Nil means all good when returned here
}
