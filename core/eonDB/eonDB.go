package eondb

import (
	"github.com/Sucks-To-Suck/Eon/helper/logger"
	"github.com/syndtr/goleveldb/leveldb"
)

// Create and return a lunch db.
// Input the name of the db save.
func NewDatabase(name string) *leveldb.DB {

	// Create the database
	db, err := leveldb.OpenFile("saves/db/"+name, nil)

	// Create a logger
	logger := logger.NewLogger("Database")

	if err != nil {

		logger.LogRed(err.Error())

	}

	logger.LogGreen("Database loaded")

	return db
}
