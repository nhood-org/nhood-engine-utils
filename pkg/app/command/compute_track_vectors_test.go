package command

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/stretchr/testify/require"
)

func Test_ComputeTrackVectorsCommandHandler_Handle(t *testing.T) {
	tracks := []model.Track{
		{
			ID:     "TCK_ID",
			Artist: "ARTIST",
			Title:  "TITLE",
			Tags: map[string]float64{
				"TAG_1": 1.0,
				"TAG_2": 0.5,
			},
		},
	}

	tags := []model.Tag{
		{
			Name:   "TAG_1",
			Vector: []float64{0.1, 1, 10, 100, 1000},
		},
		{
			Name:   "TAG_2",
			Vector: []float64{0.2, 2, 20, 200, 2000},
		},
	}

	out := new(bytes.Buffer)
	cmd := ComputeTrackVectorsCmd{
		Tracks:     tracks,
		Tags:       tags,
		VectorSize: 5,
		Output:     out,
	}

	handler := NewComputeTrackVectorsCommandHandler()

	err := handler.Handle(cmd)
	require.NoError(t, err)

	outBytes, err := ioutil.ReadAll(out)
	require.NoError(t, err)

	expectedOutput := "TCK_ID,ARTIST,TITLE,0.20000,2.00000,20.00000,200.00000,2000.00000"
	require.Equal(t, expectedOutput, string(outBytes))
}
