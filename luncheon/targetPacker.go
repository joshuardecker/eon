package luncheon

import (
	"github.com/GoblinBear/beson/types"
)

type TargetUnpacker struct {
	unpackedTarget types.UInt256 //TODO: see why has to be pointer

	packedTarget uint32
	exponent     uint32
}

func (t TargetUnpacker) Unpack(packedTarget uint32) types.UInt256 {
	// Init the t.unpackedTarget (as without the uint256 doesnt work)
	t.unpackedTarget = *types.NewUInt256("0", 1) // Temp value just used to initialize

	t.packedTarget = packedTarget

	t.exponent = t.packedTarget >> (3 * 8) // Gets the first byte out of it (big endian style)

	t.unpackedTarget.Set(Uint32toB(t.packedTarget & 0x00ffffff)) // Masks the first byte and inputs the rest into the uint235

	// TODO: Fix
	temp := t.unpackedTarget.LShift(uint(8 * (32 - t.exponent))) // To deal with the little endian that is this, shifts it left(right in big endian) to properly format
	t.unpackedTarget.Set(temp.Get())

	return t.unpackedTarget
}

func (t TargetUnpacker) UnpackAsBytes(packedTarget uint32) []byte {
	// Init the t.unpackedTarget (as without the uint256 doesnt work)
	t.unpackedTarget = *types.NewUInt256("0", 1) // Temp value just used to initialize

	t.packedTarget = packedTarget

	t.exponent = t.packedTarget >> (3 * 8) // Gets the first byte out of it (big endian style)

	t.unpackedTarget.Set(Uint32toB(t.packedTarget & 0x00ffffff)) // Masks the first byte and inputs the rest into the uint235

	// TODO: Fix
	temp := t.unpackedTarget.LShift(uint(8 * (32 - t.exponent))) // To deal with the little endian that is this, shifts it left(right in big endian) to properly format
	t.unpackedTarget.Set(temp.Get())

	return t.unpackedTarget.Get()
}
