package threader

import "github.com/Sucks-To-Suck/Eon/types/block"

// A proof algorithm agnostic Engine the threader can use to validate and verify blocks.
type Engine interface {

	// Validates the block.
	ValidateBlock(b *block.Block) error

	// Checks whether the block is valid or not.
	VerifyBlock(b *block.Block) bool
}
