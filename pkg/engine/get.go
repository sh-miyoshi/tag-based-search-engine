package engine

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func tagPermit(tag string) error {
	if tag == ":" || tag == "," {
		return fmt.Errorf("%s is not permitted as tag", tag)
	}
	return nil
}

func contains(list []string, find string) bool {
	for _, d := range list {
		if d == find {
			return true
		}
	}
	return false
}

// Get method return a list of file which matches to tags
//   If tags is empty, return all file list
func Get(tags []string, databaseFile string) ([]string, error) {
	// Validate tags
	for _, tag := range tags {
		if err := tagPermit(tag); err != nil {
			return []string{}, err
		}
	}

	// Open DB File
	fp, err := os.OpenFile(databaseFile, os.O_RDWR, 0644)
	if err != nil {
		return []string{}, err
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.Comment = '#'

	fileList := []string{}

	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				// Unexpected error occured
				return []string{}, err
			}
		}

		if len(tags) == 0 { // tags are empty
			fileList = append(fileList, line[2])
		} else {
			// append result if tags in data contains all of request tags
			dbTags := strings.Split(line[1], ":")
			matched := true
			for _, tag := range tags {
				if !contains(dbTags, tag) {
					matched = false
					break
				}
			}
			if matched {
				fileList = append(fileList, line[2])
			}
		}
	}

	return fileList, nil
}
