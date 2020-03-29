package command

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/stretchr/testify/require"
)

func Test_GenerateCorpusCommandHandler_Handle(t *testing.T) {
	out := new(bytes.Buffer)
	thresholdTag := 0.5
	thresholdSimilarity := 0.5

	cmd := GenerateCorpusCmd{
		Tracks: []model.Track{
			{
				ID: "TCK_01",
				SimilarIDs: map[string]float64{
					"TCK_02": thresholdSimilarity - 0.1,
					"TCK_03": thresholdSimilarity,
					"TCK_04": thresholdSimilarity + 0.1,
				},
				Tags: map[string]float64{
					"TAG_A": thresholdTag - 0.1,
					"TAG_B": thresholdTag,
					"TAG_C": thresholdTag + 0.1,
				},
			},
			{
				ID: "TCK_02",
				SimilarIDs: map[string]float64{
					"TCK_03": thresholdSimilarity - 0.1,
					"TCK_04": thresholdSimilarity,
					"TCK_05": thresholdSimilarity + 0.1,
				},
				Tags: map[string]float64{
					"TAG_B": thresholdTag - 0.1,
					"TAG_C": thresholdTag,
					"TAG_D": thresholdTag + 0.1,
				},
			},
		},
		ThresholdTag:        thresholdTag,
		ThresholdSimilarity: thresholdSimilarity,
		Corpus:              out,
	}

	handler := NewGenerateCorpusCommandHandler()

	err := handler.Handle(cmd)
	require.NoError(t, err)

	outBytes, err := ioutil.ReadAll(out)
	require.NoError(t, err)

	tagsCaptured := make(map[string]bool)

	outString := string(outBytes)
	for _, s := range strings.Split(outString, "\n") {
		if s == "" {
			continue
		}

		require.NotContains(t, tagsCaptured, s)
		tagsCaptured[s] = true
	}

	require.Len(t, tagsCaptured, 4)
	require.Contains(t, tagsCaptured, "TCK_01 similar to TCK_04")
	require.Contains(t, tagsCaptured, "TCK_01 tagged as TAG_C")
	require.Contains(t, tagsCaptured, "TCK_02 similar to TCK_05")
	require.Contains(t, tagsCaptured, "TCK_02 tagged as TAG_D")
}
