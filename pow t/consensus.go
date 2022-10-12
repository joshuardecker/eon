package pow

import "github.com/Sucks-To-Suck/Eon/types/block"

type WorkEngine struct{}

func (w *WorkEngine) ValidateBlock(b *block.Block) error { return nil }

func (w *WorkEngine) VerifyBlock(b *block.Block) bool { return true }
