package threader

import (
	"sync"

	poa "github.com/Sucks-To-Suck/Eon/PoA"
	pob "github.com/Sucks-To-Suck/Eon/PoB"
	pow "github.com/Sucks-To-Suck/Eon/PoW"
	"github.com/Sucks-To-Suck/Eon/core/engine"
)

type Threader struct {

	// Consensus Engine:
	engine     engine.ConsensusEngine
	engineLock *sync.WaitGroup

	// Saved engines to shift between.
	PoAEngine *poa.AuthorityEngine
	PoBEngine *pob.BurnEngine
	PoWEngine *pow.WorkEngine
}
