package target

import (
	"math/big"
)

// Unpacks the packed 32 bit target input value.
// Returns a big int that is the target.
func Unpack(packedTarget uint32) *big.Int {

	// Stores the first byte out of it (as this byte dictates the amount of bits the sig figs are from lowest value)
	exponent := packedTarget >> (3 * 8) // Gets the first byte out of it

	// Sets the non shifted value of the target
	target := big.NewInt(int64(packedTarget & 0x00ffffff))

	target.Lsh(target, (8 * uint(exponent)))

	return target
}

// Unpacks the packed 32 bit target input value.
// Returns a byte array that made up its big int form.
func UnpackAsBytes(packedTarget uint32) []byte {

	// Stores the first byte out of it (as this byte dictates the amount of bits the sig figs are from lowest value)
	exponent := packedTarget >> (3 * 8) // Gets the first byte out of it

	// Sets the non shifted value of the target
	target := big.NewInt(int64(packedTarget & 0x00ffffff))

	target.Lsh(target, (8 * uint(exponent)))

	return target.Bytes()
}
