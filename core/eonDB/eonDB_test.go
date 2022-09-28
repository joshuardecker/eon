package eondb

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {

	db := NewDatabase("dbTest")

	// Test the Put meathod.
	db.Put([]byte("1"), []byte("This is a cool test!"), nil)
	db.Put([]byte("2"), []byte("Today is the day of daying."), nil)
	db.Put([]byte("3"), []byte("Mort sends 100 to someone"), nil)

	// Test the Delete meathod.
	db.Delete([]byte("foo"), nil)
	db.Delete([]byte("1"), nil)

	// Test the Get meathod.
	data, _ := db.Get([]byte("2"), nil)
	fmt.Printf("%s\n", data) // Should print "Today is the day of daying."
}
