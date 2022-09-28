package eondb

import (
	"errors"
	"sync"
)

var (
	PUT_ERR    = errors.New("Key is already in use")
	REMOVE_ERR = errors.New("Cannot remove an element that does not exist")
	EXIST_ERR  = errors.New("Element does not exist")
	GET_ERR    = errors.New("Could not get value from the memDB")
)

// A db only stored on the memory.
// Will be used for storing data that the hard disk doesnt need to store.
type MemDb struct {
	db   map[string][]byte // Where bytes are stored
	lock sync.Mutex        // Will be used to prevent multiple go routines from accessing the db at the same time. For safety.
}

// Creates and returns a new memDb.
func InitMemDb() *MemDb {

	return &MemDb{
		db: make(map[string][]byte),
	}
}

// Adds a key and data pair to the mempool, but will not overide if the key is already used.
// Returns any errors.
func (db *MemDb) Put(key []byte, data []byte) error {

	// Lock the db so no other go routines can write while this is writting data.
	db.lock.Lock()

	// Unlock it when done.
	defer db.lock.Unlock()

	// Is the key already used?
	if db.has(key) {

		return PUT_ERR
	}

	db.db[string(key)] = data

	return nil
}

// Checks if a key already exists.
// This func is private as it does not use locking and unlocking.
// This is because only functions that already deal with locking and unlocking use this.
func (db *MemDb) has(key []byte) bool {

	// Does this key have a value that exists?
	_, exist := db.db[string(key)]

	return exist
}

// Checks if a key already exists.
// This func is public as it is protected with locks.
func (db *MemDb) Has(key []byte) bool {

	// Lock the db from being read by other go-routines.
	db.lock.Lock()

	// Unlock it at the end.
	defer db.lock.Unlock()

	// Does this key have a value that exists?
	_, exist := db.db[string(key)]

	return exist
}

// Sets the value of a key and data pair.
// Ignores if the key is already used or not.
func (db *MemDb) Set(key []byte, data []byte) error {

	// Lock the db so no other go routines can write while this is writting data.
	db.lock.Lock()

	// Unlock it when done.
	defer db.lock.Unlock()

	db.db[string(key)] = data

	return nil
}

// Deletes an element from the db.
// Gives an error if the key used does not exist in the db.
func (db *MemDb) Remove(key []byte) error {

	// Lock the db so no other go routines can write while this is writting data.
	db.lock.Lock()

	// Unlock it wen done
	defer db.lock.Unlock()

	// The the db does nothave this key
	if !db.has(key) {

		return REMOVE_ERR
	}

	delete(db.db, string(key))

	return nil
}

// Function simply gets and returns the data of a given key, no questions asked.
func (db *MemDb) get(key []byte) *[]byte {

	data := db.db[string(key)]

	return &data
}

// Gets data from the memDB based on the given key. Returns the data and any errors that occured.
func (db *MemDb) Get(key []byte) ([]byte, error) {

	// Lock the db so no other go routines can write while this is writting data.
	db.lock.Lock()

	// Unlock it wen done
	defer db.lock.Unlock()

	// Does the key exist? If not throw an error.
	if !db.has(key) {

		return nil, EXIST_ERR
	}

	// Return the data and no errors.
	return *db.get(key), nil
}
