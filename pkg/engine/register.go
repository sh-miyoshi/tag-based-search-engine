package engine

import (
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
)

// Register method register a file to database if not exists
func Register(registerFile string, databaseFile string) error {
	if registerFile == "" {
		return fmt.Errorf("register file path is empty")
	}

	// Open DB File
	fp, err := os.OpenFile(databaseFile, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comment = '#'

	// If registerFile is already exists in Database, return error
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				// Unexpected error occured
				return err
			}
		}

		if line[2] == registerFile {
			return fmt.Errorf("File is already exists")
		}
	}

	// Register File to Database
	id := uuid.New().String()
	_, err = fp.WriteString(fmt.Sprintf("%s,,%s\n", id, registerFile))

	return err
}
