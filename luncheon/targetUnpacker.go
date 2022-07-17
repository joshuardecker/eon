package luncheon

import (
	"github.com/GoblinBear/beson/types"
)

/*
This file contains the target unpacker which takes the 4 byte packed target and calculates the truncated 32 byte (256bit) target.
The way this is calculated is based off of the formula: (example here I will use 0x1d00ffff).
unpackedTarget = 0x00ffff * 2**(8 * (0x1d - 3))

This is to get it in big endian order, so some tricks were used later to make big endian but with small endian librarys
(the library for uint256 is based on little endian).
*/

// A struct that handles all the target unpacking needs.
// Can unpack targets into uint256 and byte array form.
type TargetUnpacker struct {
	unpackedTarget types.UInt256

	packedTarget uint32
	exponent     uint32

	util Util
}

// Unpacks the packed target input value.
// Returns a type uint256.
func (t *TargetUnpacker) Unpack(packedTarget uint32) types.UInt256 {

	// Init the t.unpackedTarget (as without the uint256 is blank rather than being 0's)
	t.unpackedTarget = *types.NewUInt256("0", 1)

	// Store the value in the struct for later use / organization
	t.packedTarget = packedTarget

	// Stores the first byte out of it (as this byte dictates the amount of bits the sig figs are from lowest value)
	t.exponent = t.packedTarget >> (3 * 8) // Gets the first byte out of it (big endian style)

	// Sets the non shifted value of the target
	t.unpackedTarget.Set(t.util.Uint32toB(t.packedTarget & 0x00ffffff)) // Masks the first byte and inputs the rest into the uint235

	// Swaps from little endian to big endian order
	t.byteSwap()

	// Shifts the value based on the exponent
	t.unpackedTarget = t.lShift(uint(8 * (t.exponent - 3)))

	return t.unpackedTarget
}

// Unpacks the packed target input value.
// Returns a byte array (256 bits long).
func (t *TargetUnpacker) UnpackAsBytes(packedTarget uint32) []byte {

	// Init the t.unpackedTarget (as without the uint256 is blank rather than being 0's)
	t.unpackedTarget = *types.NewUInt256("0", 1)

	// Store the value in the struct for later use / organization
	t.packedTarget = packedTarget

	// Stores the first byte out of it (as this byte dictates the amount of bits the sig figs are from lowest value)
	t.exponent = t.packedTarget >> (3 * 8) // Gets the first byte out of it (big endian style)

	// Sets the non shifted value of the target
	t.unpackedTarget.Set(t.util.Uint32toB(t.packedTarget & 0x00ffffff)) // Masks the first byte and inputs the rest into the uint235

	// Swaps from little endian to big endian order
	t.byteSwap()

	// Shifts the value based on the exponent
	t.unpackedTarget = t.lShift(uint(8 * (t.exponent - 3)))

	return t.unpackedTarget.Get()
}

// Function bitshifts left and returns the value for better looking code above.
func (t *TargetUnpacker) lShift(shiftAmount uint) types.UInt256 {

	shiftedTarget := t.unpackedTarget.RShift(shiftAmount) // This is because the uint256 library is built on little endian, so bitshifts are opposite

	return *shiftedTarget
}

// Function bitshifts right and returns the value for better looking code above.
func (t *TargetUnpacker) rShift(shiftAmount uint) types.UInt256 {

	shiftedTarget := t.unpackedTarget.LShift(shiftAmount) // This is because the uint256 library is built on little endian, so bitshifts are opposite

	return *shiftedTarget
}

// Swaps the order of the bytes in the unpacked target. Updates the value automatically. Returns nothing.
func (t *TargetUnpacker) byteSwap() {

	swapBytes := t.unpackedTarget.Get()
	t.util.ByteArraySwap(swapBytes)

	t.unpackedTarget.Set(swapBytes)
}
