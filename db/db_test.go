package db

import (
	"testing"
)

func TestOpen(t *testing.T) {
	db, err := Open()
	defer db.Close()

	if err != nil {
		t.Errorf("Could not connect to database server.")
	}
}
