package thread

import (
	"bytes"
	"errors"
	"math/big"

	eondb "github.com/Sucks-To-Suck/Eon/core/eonDB"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
	"github.com/Sucks-To-Suck/Eon/types/block"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	HASH_ERR   = errors.New("Couldnt get the hash of the block")
	DECODE_ERR = errors.New("Couldnt decode the block")
	DB_ERR     = errors.New("Database error")
	MEMDB_ERR  = errors.New("MemDatabase error")
	BUFF_ERR   = errors.New("Buffer error: ")
)

// A thread is a single thread of blocks, aka a normal blockchain.
type Thread struct {

	// The id of the thread.
	Id *big.Int

	// Databases of the thread.
	DB    *leveldb.DB
	memDB *eondb.MemDb
}

// Creates and returns a new thread with everything generated, ready for use.
func NewThread(Id *big.Int) *Thread {

	return &Thread{
		Id:    Id,
		DB:    eondb.NewDB(Id.String()),
		memDB: eondb.InitMemDb(),
	}
}

// Update the id of the thread.
func (t *Thread) ChangeId(b big.Int) {

	t.Id = &b
}

// ****
// DB:

// Adds a block to the DB.
func (t *Thread) AddBlock(b *block.Block) error {

	// Get the hash of the block.
	bHash := b.GetHash()

	// Store it in the db.
	return t.DB.Put(bHash.GetBytes(), b.Bytes(), nil)
}

// Removes the given block from the DB.
func (t *Thread) RemoveBlock(b *block.Block) error {

	// Get the hash of the block.
	bHash := b.GetHash()

	return t.DB.Delete(bHash.GetBytes(), nil)
}

// Removes a block based on its hash from the DB.
func (t *Thread) RemoveBlockByHash(h *eocrypt.Hash) error {

	return t.DB.Delete(h.GetBytes(), nil)
}

// Gets and returns the block based on the given hash from the DB.
func (t *Thread) GetBlockByHash(h *eocrypt.Hash) (b *block.Block, e error) {

	// Get the block as bytes from the DB.
	blockBytes, e := t.DB.Get(h.GetBytes(), nil)

	// If the block couldnt be got by the DB.
	if e != nil {

		return nil, DB_ERR
	}

	// Create a bytes buffer for the block bytes.
	buff := new(bytes.Buffer)

	// Write the block into the bytes buffer.
	_, e = buff.Write(blockBytes)

	// If the block couldnt be written into the buffer.
	if e != nil {

		return nil, BUFF_ERR
	}

	// Decode the block bytes into a block.
	b, e = block.Decode(buff)

	// If the block could not be decoded.
	if e != nil {

		return nil, DECODE_ERR
	}

	return b, nil
}

// DB:
// ****

// ****
// MemDB:

// Adds a block to the memoryDB.
func (t *Thread) AddBlockMem(b *block.Block) error {

	// Get the hash of the block.
	bHash := b.GetHash()

	return t.memDB.Set(bHash.GetBytes(), b.Bytes())
}

// Removes the given block from the memoryDB.
func (t *Thread) RemoveBlockMem(b *block.Block) error {

	// Get the hash of the block.
	bHash := b.GetHash()

	return t.memDB.Remove(bHash.GetBytes())
}

// Removes the given block by the block hash given.
func (t *Thread) RemoveBlockByHashMem(h *eocrypt.Hash) error {

	return t.memDB.Remove(h.GetBytes())
}

// Gets and returns the block based on the given hash from the memDB.
func (t *Thread) GetBlockByHashMem(h *eocrypt.Hash) (b *block.Block, e error) {

	// Get the block as bytes from the memDB.
	blockBytes, e := t.memDB.Get(h.GetBytes())

	// If the block couldnt be got by the memDB.
	if e != nil {

		return nil, MEMDB_ERR
	}

	// Create a bytes buffer for the block bytes.
	buff := new(bytes.Buffer)

	// Write the block into the bytes buffer.
	_, e = buff.Write(blockBytes)

	// If the block couldnt be written into the buffer.
	if e != nil {

		return nil, BUFF_ERR
	}

	// Decode the block bytes into a block.
	b, e = block.Decode(buff)

	// If the block could not be decoded.
	if e != nil {

		return nil, DECODE_ERR
	}

	return b, nil
}

// MemDB:
// ****
