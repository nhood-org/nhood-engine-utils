package model

import (
	"encoding/json"
	"os"
)

/*
Track defines an internal track entity

*/
type Track struct {
	ID         string   `json:"track_id"`
	SimilarIDs []string `json:"similar_ids"`
	Tags       []string `json:"tags"`
}

/*
Tracks is a map of track entities with its ID as a key

*/
type Tracks map[string]Track

/*
ReadTracksFromFile reads tracks from given json file

*/
func ReadTracksFromFile(fileName string) (*Tracks, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer finalize(file)

	decoder := json.NewDecoder(file)
	t := Tracks{}
	err = decoder.Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
