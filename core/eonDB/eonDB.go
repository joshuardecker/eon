package eondb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

// Create / open a DB with the given file name.
func NewDB(name string) *leveldb.DB {

	// Create the database.
	db, dbErr := leveldb.OpenFile("saves/db/"+name, nil)

	// If the db is corrupted and needs to be recovered.
	if errors.IsCorrupted(dbErr) {

		db, dbErr = leveldb.RecoverFile("saves/db/"+name, nil)

	}

	// Panic if the DB could not be loaded and recovering didnt work / wasnt an option.
	if dbErr != nil {

		panic("DB could not be loaded, Error: " + dbErr.Error())
	}

	return db
}
