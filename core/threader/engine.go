package threader

import "errors"

// The types of engines available to run on, represented by ints.
const (
	PoA = 1
	PoB = 2
	PoW = 3
)

var (
	SHIFTERROR = errors.New("Could not shift engines given the gear.")
)

// Shift the consensus engine to a different consensus mechanism.
func (t *Threader) Shift(gear int) error {

	// Waits for the engine to not be in use before shifting.
	t.engineLock.Wait()

	switch gear {

	// Switch to PoA?
	case PoA:
		t.engine = t.PoAEngine

		return nil

	// Switch to PoB?
	case PoB:
		t.engine = t.PoBEngine

		return nil

	// Switch to PoB.
	case PoW:
		t.engine = t.PoWEngine

		return nil

	// Couldnt switch engines.
	default:
		return SHIFTERROR
	}
}
