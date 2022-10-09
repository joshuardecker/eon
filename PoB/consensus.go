package pob

import "github.com/Sucks-To-Suck/Eon/types/block"

type BurnEngine struct{}

func (b *BurnEngine) ValidateBlock(bl *block.Block) error { return nil }

func (b *BurnEngine) VerifyBlock(bl *block.Block) bool { return true }
