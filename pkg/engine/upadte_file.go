package engine

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
)

// RegisterFile method register a file to database if not exists
func RegisterFile(registerFile string, databaseFile string) error {
	if registerFile == "" {
		return fmt.Errorf("register file path is empty")
	}

	// Open DB File
	raw, err := ioutil.ReadFile(databaseFile)
	if err != nil {
		return err
	}

	var data []object
	json.Unmarshal(raw, &data)

	for _, d := range data {
		if registerFile == d.FilePath {
			return fmt.Errorf("File is already exists")
		}
	}

	data = append(data, object{
		ID:       uuid.New().String(),
		Tags:     []string{},
		FilePath: registerFile,
	})

	// Update file
	updated, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(databaseFile, updated, 0644)
	return nil
}

// DeleteFile method delete a file to database if not exists
func DeleteFile(deleteFile string, databaseFile string) error {
	// TODO(Delete File from database)
	return nil
}
