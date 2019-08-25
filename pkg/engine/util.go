package engine

import (
	"fmt"
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
