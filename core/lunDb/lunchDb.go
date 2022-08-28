package lundb

import (
	"github.com/Sucks-To-Suck/LuncheonNetwork/logger"
	"github.com/syndtr/goleveldb/leveldb"
)

// Create and return a lunch db.
// Input the name of the db save.
func InitDatabase(name string) *leveldb.DB {

	// Create the database
	db, err := leveldb.OpenFile("saves/db/"+name, nil)

	// Create a logger
	logger := logger.NewLogger("Database")

	if err != nil {

		logger.LogRed(err.Error())

	}

	logger.LogGreen("database loaded")

	return db
}
