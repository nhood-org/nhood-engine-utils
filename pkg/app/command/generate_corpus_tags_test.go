package command

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/stretchr/testify/require"
)

func Test_GenerateCorpusTagsCommandHandler_Handle(t *testing.T) {
	out := new(bytes.Buffer)
	cmd := GenerateCorpusTagsCmd{
		Tracks: []model.Track{
			{
				Tags: map[string]float64{
					"TAG_A": 1.0,
					"TAG_B": 1.0,
					"TAG_C": 1.0,
				},
			},
			{
				Tags: map[string]float64{
					"TAG_B": 1.0,
					"TAG_C": 1.0,
					"TAG_D": 1.0,
				},
			},
		},
		Corpus: out,
	}

	handler := NewGenerateCorpusTagsCommandHandler()

	err := handler.Handle(cmd)
	require.NoError(t, err)

	outBytes, err := ioutil.ReadAll(out)
	require.NoError(t, err)

	outString := string(outBytes)
	outStringLines := strings.Split(outString, "\n")
	require.Len(t, outStringLines, 3)
	require.Empty(t, outStringLines[2])

	lineTagsCaptured := make(map[string]bool)
	line := strings.Split(outStringLines[0], " ")
	for _, tag := range line {
		_, exists := lineTagsCaptured[tag]
		if exists {
			require.FailNowf(t, "tag %s duplicated in line %s", tag, lineTagsCaptured)
		}

		lineTagsCaptured[tag] = true
	}

	require.Len(t, lineTagsCaptured, 3)
	require.Contains(t, lineTagsCaptured, "TAG_A")
	require.Contains(t, lineTagsCaptured, "TAG_B")
	require.Contains(t, lineTagsCaptured, "TAG_C")

	lineTagsCaptured = make(map[string]bool)
	line = strings.Split(outStringLines[1], " ")
	for _, tag := range line {
		_, exists := lineTagsCaptured[tag]
		if exists {
			require.FailNowf(t, "tag %s duplicated in line %s", tag, lineTagsCaptured)
		}

		lineTagsCaptured[tag] = true
	}

	require.Len(t, lineTagsCaptured, 3)
	require.Contains(t, lineTagsCaptured, "TAG_B")
	require.Contains(t, lineTagsCaptured, "TAG_C")
	require.Contains(t, lineTagsCaptured, "TAG_D")
}
