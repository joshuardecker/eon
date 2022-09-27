package thread

import (
	"math/big"

	eondb "github.com/Sucks-To-Suck/Eon/core/eonDB"
	"github.com/syndtr/goleveldb/leveldb"
)

// A thread is a single thread of blocks, aka a normal blockchain.
type Thread struct {

	// The id of the thread.
	Id *big.Int

	// Databases of the thread.
	DB    *leveldb.DB
	memDB *eondb.MemDb
}
