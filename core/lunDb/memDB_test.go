package lundb

import (
	"fmt"
	"testing"
)

func TestMemDB(t *testing.T) {

	mDB := InitMemDb()

	// Data
	data1 := []byte("This is pretty cool!")
	data2 := []byte("1.2.3")
	data3 := []byte("This is a test.")

	// Test the Put meathod
	mDB.Put([]byte("1"), data1)
	mDB.Put([]byte("2"), data2)
	mDB.Put([]byte("3"), data3)
	fmt.Printf("mDB: %v\n", mDB)

	// Test the Remove Meathod
	mDB.Remove([]byte("1"))
	fmt.Printf("mDB: %v\n", mDB)

	// Test the has meathod
	fmt.Print(mDB.Has([]byte("1")), "\n") // Should print false
	fmt.Print(mDB.Has([]byte("2")), "\n") // Should print true

	// Test the Set meathod
	fmt.Printf("mDB: %v\n", mDB)
	mDB.Set([]byte("2"), []byte("Partif"))
	fmt.Printf("mDB: %v\n", mDB) // Should not print the same data, but same keys
}
