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

	tagsCaptured := make(map[string]bool)

	outString := string(outBytes)
	for _, s := range strings.Split(outString, " ") {
		if s == "" {
			continue
		}

		require.NotContains(t, tagsCaptured, s)
		tagsCaptured[s] = true
	}

	require.Len(t, tagsCaptured, 4)
	require.Contains(t, tagsCaptured, "TAG_A")
	require.Contains(t, tagsCaptured, "TAG_B")
	require.Contains(t, tagsCaptured, "TAG_C")
	require.Contains(t, tagsCaptured, "TAG_D")
}
