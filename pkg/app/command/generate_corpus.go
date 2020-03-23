package command

import (
	"fmt"
	"io"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

type GenerateCorpusCmd struct {
	Tracks              []model.Track
	ThresholdTag        float64
	ThresholdSimilarity float64
	Corpus              io.Writer
}

type GenerateCorpusCommandHandler interface {
	Handle(GenerateCorpusCmd) error
}

type generateCorpusCommandHandler struct {
}

func NewGenerateCorpusCommandHandler() GenerateCorpusCommandHandler {
	return generateCorpusCommandHandler{}
}

func (h generateCorpusCommandHandler) Handle(cmd GenerateCorpusCmd) error {
	for _, t := range cmd.Tracks {
		for id, weight := range t.SimilarIDs {
			if weight > cmd.ThresholdSimilarity {
				appendSimilarID(cmd.Corpus, t.ID, id)
			}
		}
		for tag, weight := range t.Tags {
			if weight > cmd.ThresholdTag {
				appendTag(cmd.Corpus, t.ID, tag)
			}
		}
	}
	return nil
}

func appendSimilarID(writer io.Writer, trackID string, id string) {
	fmt.Fprintf(writer, "%s similar to %s\n", trackID, id)
}

func appendTag(writer io.Writer, trackID string, tag string) {
	fmt.Fprintf(writer, "%s tagged as %s\n", trackID, tag)
}
