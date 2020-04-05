package command

import (
	"fmt"
	"io"
	"strings"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/pkg/errors"
)

type ComputeTrackVectorsCmd struct {
	Tracks     []model.Track
	Tags       []model.Tag
	VectorSize int
	Output     io.Writer
}

type ComputeTrackVectorsCommandHandler interface {
	Handle(ComputeTrackVectorsCmd) error
}

type computeTrackVectorsCommandHandler struct {
}

func NewComputeTrackVectorsCommandHandler() ComputeTrackVectorsCommandHandler {
	return computeTrackVectorsCommandHandler{}
}

func (h computeTrackVectorsCommandHandler) Handle(cmd ComputeTrackVectorsCmd) error {
	tags, err := mapTagVectorsById(cmd.Tags, cmd.VectorSize)
	if err != nil {
		return err
	}

	trackVectors := make(map[string][]float64)
	for _, t := range cmd.Tracks {
		trackVectors[t.ID] = computeTrackVector(t, tags, cmd.VectorSize)
	}

	renderTracksToWriter(cmd.Tracks, trackVectors, cmd.Output)

	return nil
}

func mapTagVectorsById(
	tags []model.Tag,
	expectedSize int,
) (map[string][]float64, error) {
	m := make(map[string][]float64)
	for _, t := range tags {
		vSize := len(t.Vector)
		if vSize != expectedSize {
			return nil, errors.Errorf(
				"tag '%s' has invalid vector size %d, expected is %d", t.Name, vSize, expectedSize)
		}
		m[t.Name] = t.Vector
	}
	return m, nil
}

func computeTrackVector(
	track model.Track,
	tags map[string][]float64,
	expectedSize int,
) []float64 {
	vector := make([]float64, expectedSize)
	for tag, weight := range track.Tags {
		tagVector, ok := tags[tag]
		if !ok {
			continue
		}
		for i, v := range tagVector {
			vector[i] += v * weight
		}
	}
	return vector
}

func renderTracksToWriter(
	tracks []model.Track,
	trackVectors map[string][]float64,
	output io.Writer,
) {
	for _, t := range tracks {
		trackVector := trackVectors[t.ID]
		fmt.Fprintf(output, "%s\n", renderTrack(t, trackVector))
	}
}

func renderTrack(
	track model.Track,
	trackVector []float64,
) string {
	trackStrings := []string{
		track.ID,
		track.Artist,
		track.Title,
		renderTrackVector(trackVector),
	}
	return strings.Join(trackStrings, ",")
}

func renderTrackVector(
	vector []float64,
) string {
	trackVectorStrings := make([]string, len(vector))
	for i, v := range vector {
		trackVectorStrings[i] = fmt.Sprintf("%.10f", v)
	}
	return strings.Join(trackVectorStrings, ",")
}
