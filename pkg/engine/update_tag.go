package engine

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// Assign method assign tags to targetFile
func Assign(targetFile string, tags []string, databaseFile string) error {
	if targetFile == "" {
		return fmt.Errorf("target file path is empty")
	}

	// Open DB File
	fp, err := os.OpenFile(databaseFile, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comment = '#'

	var records [][]string

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

		if line[2] == targetFile {
			// TODO(udpate tags)
		}

		records = append(records, line)
	}

	// Remove All data at first
	fp.Truncate(0)
	fp.Seek(0, 0)

	// Write updated data
	writer := csv.NewWriter(fp)
	writer.WriteAll(records)

	return err
}

// Delete method delete tags from targetFile
func Delete(targetFile string, tags []string, databaseFile string) error {
	// TODO
	return nil
}
