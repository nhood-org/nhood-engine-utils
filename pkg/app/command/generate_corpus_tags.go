package command

import (
	"fmt"
	"io"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

type GenerateCorpusTagsCmd struct {
	Tracks []model.Track
	Corpus io.Writer
}

type GenerateCorpusTagsCommandHandler interface {
	Handle(GenerateCorpusTagsCmd) error
}

type generateCorpusTagsCommandHandler struct {
}

func NewGenerateCorpusTagsCommandHandler() GenerateCorpusTagsCommandHandler {
	return generateCorpusTagsCommandHandler{}
}

func (h generateCorpusTagsCommandHandler) Handle(cmd GenerateCorpusTagsCmd) error {
	tags := make(map[string]bool)

	for _, t := range cmd.Tracks {
		for tag := range t.Tags {
			if _, ok := tags[tag]; !ok {
				tags[tag] = true
			}
		}
	}

	for tag := range tags {
		fmt.Fprintf(cmd.Corpus, "%s ", tag)
	}

	return nil
}
