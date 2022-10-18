package thread

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/Sucks-To-Suck/Eon/core/gas"
	"github.com/Sucks-To-Suck/Eon/eocrypt"
	"github.com/Sucks-To-Suck/Eon/tools/eondb"
	"github.com/Sucks-To-Suck/Eon/types/block"
	"github.com/Sucks-To-Suck/Eon/types/transaction"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	HASH_ERR   = errors.New("Couldnt get the hash of the block")
	DECODE_ERR = errors.New("Couldnt decode the block")
	DB_ERR     = errors.New("Database error")
	MEMDB_ERR  = errors.New("MemDatabase error")
	BUFF_ERR   = errors.New("Buffer error")

	RECENT_BLOCK = []byte("Most Recent Block") // The byte key used to get the most recent block.
)

// A thread is a single thread of blocks, aka a normal blockchain.
type Thread struct {

	// Databases of the thread.
	DB    *leveldb.DB
	memDB *eondb.MemDb
	txs   *[]transaction.Transaction
}

// Creates and returns a new thread with everything generated, ready for use.
// Uses a Big Int as the Id and not a uint64 because a Big Int can be accurately be converted to a string, which is needed for the DB save name.
func NewThread(Id *big.Int) *Thread {

	return &Thread{
		DB:    eondb.NewDB(Id.String()),
		memDB: eondb.NewMemDb(),
	}
}

// ****
// DB:

// Adds a block to the DB.
func (t *Thread) AddBlock(b *block.Block) error {

	// Get the hash of the block.
	bHash := b.Hash()

	// Store this block as the most recent block in the DB.
	recentBlockErr := t.DB.Put(RECENT_BLOCK, b.Bytes(), nil)

	// If an error occured.
	if recentBlockErr != nil {

		return recentBlockErr
	}

	// Store it normally in the db as well.
	return t.DB.Put(bHash.Bytes(), b.Bytes(), nil)
}

// Removes the given block from the DB.
func (t *Thread) RemoveBlock(b *block.Block) error {

	// Get the hash of the block.
	bHash := b.Hash()

	return t.DB.Delete(bHash.Bytes(), nil)
}

// Removes a block based on its hash from the DB.
func (t *Thread) RemoveBlockByHash(h *eocrypt.Hash) error {

	return t.DB.Delete(h.Bytes(), nil)
}

// Gets and returns the block based on the given hash from the DB.
func (t *Thread) GetBlockByHash(h *eocrypt.Hash) (b *block.Block, e error) {

	// Get the block as bytes from the DB.
	blockBytes, e := t.DB.Get(h.Bytes(), nil)

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

// Gets and returns the most recent block from the DB.
func (t *Thread) GetRecentBlock() (b *block.Block, e error) {

	// Get the block as bytes from the DB.
	blockBytes, e := t.DB.Get(RECENT_BLOCK, nil)

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
	bHash := b.Hash()

	return t.memDB.Set(bHash.Bytes(), b.Bytes())
}

// Removes the given block from the memoryDB.
func (t *Thread) RemoveBlockMem(b *block.Block) error {

	// Get the hash of the block.
	bHash := b.Hash()

	return t.memDB.Remove(bHash.Bytes())
}

// Removes the given block by the block hash given.
func (t *Thread) RemoveBlockByHashMem(h *eocrypt.Hash) error {

	return t.memDB.Remove(h.Bytes())
}

// Gets and returns the block based on the given hash from the memDB.
func (t *Thread) GetBlockByHashMem(h *eocrypt.Hash) (b *block.Block, e error) {

	// Get the block as bytes from the memDB.
	blockBytes, e := t.memDB.Get(h.Bytes())

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

// ****
// Transactions:

// Add the tx to the pending tx list.
func (t *Thread) AddTx(tx *transaction.Transaction) {

	*t.txs = append(*t.txs, *tx)
}

// Retrieves pending transactions until either no transactions are pending or the gas limit for the next block is hit.
func (t *Thread) RetrieveTxs(gasLim gas.Gas) (finalTxs *[]transaction.Transaction) {

	// If there are no txs pending.
	if len(*t.txs) == 0 {

		return nil
	}

	// Loop through adding txs to the finalTxs until gasLimit hits 0, aka the gas limit has been hit.
	for _, tx := range *t.txs {

		// Add the tx to the final tx slice.
		*finalTxs = append(*finalTxs, tx)

		// Subtract the tx gas from the gasLimit.
		gasLim -= tx.GetGas()

		// If the limit has been reached.
		if gasLim == 0 {

			return
		}
	}

	return
}

// Transactions:
// ****
