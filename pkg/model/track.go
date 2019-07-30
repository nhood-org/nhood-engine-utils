package model

/*
Track is a song entity

*/
type Track struct {
	ID     string `json:"track_id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

/*
ToString returns a string representation of a track entity

*/
func (t *Track) ToString() string {
	return t.ID + ": " + t.Artist + " - " + t.Title
}
