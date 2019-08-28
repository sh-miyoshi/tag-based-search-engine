package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// AssignTag method assign tags to targetFile
func AssignTag(targetFile string, tags []string, databaseFile string) error {
	if targetFile == "" {
		return fmt.Errorf("target file path is empty")
	}

	// Open DB File
	raw, err := ioutil.ReadFile(databaseFile)
	if err != nil {
		return err
	}

	var data []object
	json.Unmarshal(raw, &data)

	for index, d := range data {
		if d.FilePath == targetFile {
			// Update Tags
			for _, tag := range tags {
				if !contains(data[index].Tags, tag) {
					data[index].Tags = append(data[index].Tags, tag)
				}
			}
			break
		}
	}

	// Update file
	updated, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(databaseFile, updated, 0644)
	return nil
}

// DeleteTag method delete tags from targetFile
func DeleteTag(targetFile string, tags []string, databaseFile string) error {
	if targetFile == "" {
		return fmt.Errorf("target file path is empty")
	}

	// Open DB File
	raw, err := ioutil.ReadFile(databaseFile)
	if err != nil {
		return err
	}

	var data []object
	json.Unmarshal(raw, &data)

	for index, d := range data {
		if d.FilePath == targetFile {
			// Delete Tags
			resTags := []string{}
			for _, tag := range data[index].Tags {
				if !contains(tags, tag) {
					resTags = append(resTags, tag)
				}
			}
			data[index].Tags = append([]string{}, resTags...)
			break
		}
	}

	// Update file
	updated, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(databaseFile, updated, 0644)
	return nil
}
