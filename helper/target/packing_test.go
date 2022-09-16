package target

import (
	"fmt"
	"testing"
)

func TestTarget(t *testing.T) {

	// Packed target
	pTarget1 := uint32(0x15ffffff)
	pTarget2 := uint32(0x1cf00fff)
	pTarget3 := uint32(0x0856def3)

	target1 := Unpack(pTarget1)
	target2 := Unpack(pTarget2)
	target3 := Unpack(pTarget3)

	fmt.Printf("Target 1:%x\n", target1)
	fmt.Printf("Target 2:%x\n", target2)
	fmt.Printf("Target 3:%x\n", target3)

	if pTarget1 != PackTargetInt(target1) {

		fmt.Println("Target 1 did not work!")
		fmt.Printf("Target 1:%x\n", PackTargetInt(target1))

	}

	if pTarget2 != PackTargetInt(target2) {

		fmt.Println("Target 2 did not work!")
		fmt.Printf("Target 2:%x\n", PackTargetInt(target2))

	}

	if pTarget3 != PackTargetInt(target3) {

		fmt.Println("Target 3 did not work!")
		fmt.Printf("Target 2:%x\n", PackTargetInt(target3))

	}
}
