package model

import (
	"encoding/json"
	"os"
)

/*
Tags is a slice of track tag entities

*/
type Tags map[string]bool

/*
ReadTagsFromFile reads track tags from given json file

*/
func ReadTagsFromFile(fileName string) (*Tags, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer finalize(file)

	decoder := json.NewDecoder(file)
	t := Tags{}
	err = decoder.Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

/*
Contains checks if tag of tagName exists

*/
func (t *Tags) Contains(tagName string) bool {
	_, contains := (*t)[tagName]
	return contains
}
