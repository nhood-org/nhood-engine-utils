package model

import (
	"encoding/json"
	"os"
)

type RawTrack struct {
	ID         string          `json:"track_id"`
	SimilarIDs [][]interface{} `json:"similars"`
	Artist     string          `json:"artist"`
	Title      string          `json:"title"`
	Tags       [][]string      `json:"tags"`
}

func ReadRawTrackFromFile(fileName string) (*RawTrack, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer finalize(file)

	decoder := json.NewDecoder(file)
	c := RawTrack{}
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func finalize(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}
