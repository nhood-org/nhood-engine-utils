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
			Vector: []float64{1, 0.1, 0.01, 0.001, 0.0001},
		},
		{
			Name:   "TAG_2",
			Vector: []float64{2, 0.2, 0.02, 0.002, 0.0002},
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

	expectedOutput := "TCK_ID,ARTIST,TITLE,2.0000000000,0.2000000000,0.0200000000,0.0020000000,0.0002000000\n"
	require.Equal(t, expectedOutput, string(outBytes))
}
