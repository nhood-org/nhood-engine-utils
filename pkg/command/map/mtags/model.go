package mtags

import "github.com/nhood-org/nhood-engine-utils/pkg/utils"

/*
tag is a song tag entity

*/
type tag struct {
	Name   string
	Weight int64
}

/*
tagStatistics defines a tag and its statistics

*/
type tagStatistics struct {
	Name       string
	Statistics *utils.MovingStatistics
}
