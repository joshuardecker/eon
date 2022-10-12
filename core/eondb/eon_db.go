package eondb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

var (
	DB_ERR = errors.New("DB could not be loaded: ")
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

		panic(DB_ERR.Error() + dbErr.Error())
	}

	return db
}
