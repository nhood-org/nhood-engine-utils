package gvec

import (
	"log"

	"github.com/nhood-org/nhood-engine-utils/pkg/command/gen/gvec/relevance"
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

type MapperImpl struct {
}

type MapperImplContext struct {
	tags   *model.Tags
	tracks *model.Tracks
	matrix *relevance.Matrix
}

func NewMapperImpl() *MapperImpl {
	return &MapperImpl{}
}

func (t *MapperImpl) resolve(tags *model.Tags, tracks *model.Tracks) *relevance.Matrix {
	ctx := MapperImplContext{
		tags:   tags,
		tracks: tracks,
		matrix: relevance.NewMatrix(tags),
	}
	ctx.resolve()
	return ctx.matrix
}

func (t *MapperImplContext) resolve() {
	i := 0
	for _, track := range *t.tracks {
		t.mapTagRelevancyWithinTrack(&track)
		if i%1000 == 0 {
			log.Println(i, "tracks mapped")
		}
		i++
	}
}

func (t *MapperImplContext) mapTagRelevancyWithinTrack(track1 *model.Track) {
	for _, ID := range track1.SimilarIDs {
		track2 := (*t.tracks)[ID]
		t.mapTagRelevancyRelevanceOfTracks(track1, &track2)
	}
}

func (t *MapperImplContext) mapTagRelevancyRelevanceOfTracks(track1 *model.Track, track2 *model.Track) {
	for _, t1 := range track1.Tags {
		t.mapTagRelevancyRelevanceOfTags(t1, track2.Tags)
	}
}

func (t *MapperImplContext) mapTagRelevancyRelevanceOfTags(t1 string, tags []string) {
	for _, t2 := range tags {
		t.matrix.Increment(t1, t2)
	}
}
