package model

import (
	"encoding/json"
	"os"
)

/*
Track is a song entity

*/
type Track struct {
	ID     string `json:"track_id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

/*
ReadTrack reads configuration from given json file

*/
func ReadTrack(fileName string) (*Track, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	c := Track{}
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

/*
ToString return a string representation of a track entity

*/
func (t *Track) ToString() string {
	return t.ID + ": " + t.Artist + " - " + t.Title
}
