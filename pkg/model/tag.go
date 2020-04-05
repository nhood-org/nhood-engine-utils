package model

type Tag struct {
	Name   string    `json:"name"`
	Vector []float64 `json:"vector"`
}
