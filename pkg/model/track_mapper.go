package model

import (
	"encoding/json"
	"os"
)

/*
ReadTrack reads configuration from given json file

*/
func ReadTrack(fileName string) (*Track, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer finalizeTrackReading(file)

	decoder := json.NewDecoder(file)
	c := Track{}
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func finalizeTrackReading(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}
