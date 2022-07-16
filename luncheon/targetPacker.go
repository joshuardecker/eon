package luncheon

import (
	"encoding/binary"

	"github.com/GoblinBear/beson/types"
)

type TargetPacker struct {
	unpackedTarget      types.UInt256
	unpackedTargetBytes []byte

	packedTarget uint32
	exponent     uint32

	util Util
}

func (t TargetPacker) PackTargetUint256(unpackedTarget types.UInt256) uint32 {

	t.packedTarget = 0

	t.unpackedTarget = unpackedTarget
	t.unpackedTargetBytes = t.unpackedTarget.Get()

	var a uint32
	for a = 0; t.unpackedTargetBytes[a] == byte(0); a++ {
	}

	t.exponent = 32 - a

	byteArray := make([]byte, 4)

	byteArray[0] = byte(t.exponent)
	byteArray[1] = t.unpackedTargetBytes[a]
	byteArray[2] = t.unpackedTargetBytes[a+1]
	byteArray[3] = t.unpackedTargetBytes[a+2]

	t.packedTarget = binary.BigEndian.Uint32(byteArray)

	return t.packedTarget
}

func (t TargetPacker) PackTargetBytes(unpackedTargetBytes []byte) uint32 {

	t.packedTarget = 0

	t.unpackedTarget.Set(unpackedTargetBytes)
	t.unpackedTargetBytes = unpackedTargetBytes

	var a uint32
	for a = 0; t.unpackedTargetBytes[a] == byte(0); a++ {

		if a == 28 {

			break
		}
	}

	t.exponent = 32 - a

	byteArray := make([]byte, 4)

	byteArray[0] = byte(t.exponent)
	byteArray[1] = t.unpackedTargetBytes[a]
	byteArray[2] = t.unpackedTargetBytes[a+1]
	byteArray[3] = t.unpackedTargetBytes[a+2]

	t.packedTarget = binary.BigEndian.Uint32(byteArray)

	return t.packedTarget
}
