package engine

import (
	"encoding/json"
	"io/ioutil"
)

// Get method return a list of file which matches to tags
//   If tags is empty, return all file list
func Get(tags []string, databaseFile string) ([]string, error) {
	// Open DB File
	raw, err := ioutil.ReadFile(databaseFile)
	if err != nil {
		return []string{}, err
	}

	var data []object
	json.Unmarshal(raw, &data)

	fileList := []string{}

	for _, d := range data {
		if len(tags) == 0 { // tags are empty
			fileList = append(fileList, d.FilePath)
		} else {
			// append result if tags in data contains all of request tags
			matched := true
			for _, tag := range tags {
				if !contains(d.Tags, tag) {
					matched = false
					break
				}
			}
			if matched {
				fileList = append(fileList, d.FilePath)
			}
		}
	}

	return fileList, nil
}
