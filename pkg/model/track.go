package model

/*
Track is a song entity

*/
type Track struct {
	ID     string     `json:"track_id"`
	Artist string     `json:"artist"`
	Title  string     `json:"title"`
	Tags   [][]string `json:"tags"`
}

/*
TrackTag is a song tag entity

*/
type TrackTag struct {
	Name   string
	Weight int64
}
