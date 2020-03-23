package model

import "strconv"

type Track struct {
	ID         string             `json:"track_id"`
	SimilarIDs map[string]float64 `json:"similar_ids"`
	Artist     string             `json:"artist"`
	Title      string             `json:"title"`
	Tags       map[string]float64 `json:"tags"`
}

func TrackFromRaw(raw RawTrack) Track {
	similarIDs := make(map[string]float64)
	for _, id := range raw.SimilarIDs {
		if len(id) == 2 {
			similarIDs[id[0].(string)] = id[1].(float64)
		}
	}

	tags := make(map[string]float64)
	for _, tag := range raw.Tags {
		if len(tag) == 2 {
			weight, _ := strconv.Atoi(tag[1])
			tags[tag[0]] = float64(weight)
		}
	}

	return Track{
		ID:         raw.ID,
		SimilarIDs: similarIDs,
		Artist:     raw.Artist,
		Title:      raw.Title,
		Tags:       tags,
	}
}
