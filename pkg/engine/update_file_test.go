package engine

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRegisterFile(t *testing.T) {
	dbFile := ""
	if err := RegisterFile("regiter_file.txt", dbFile); err == nil {
		t.Errorf("Expect error when empty database file, but got nil")
	}

	dbFile = "unexists_db.json"
	if err := RegisterFile("regiter_file.txt", dbFile); err == nil {
		t.Errorf("Expect error when read unexists database file, but got nil")
	}

	tmpfile, _ := ioutil.TempFile("", "")
	defer os.Remove(tmpfile.Name())
	dbFile = tmpfile.Name()

	if err := RegisterFile("regiter_file.txt", dbFile); err != nil {
		t.Errorf("Try to register new file, but got error: %v", err)
	}

	// Try to register duplicate file
	if err := RegisterFile("regiter_file.txt", dbFile); err == nil {
		t.Errorf("Expect error when registering duplicate file, but got nil")
	}
}
