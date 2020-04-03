package command

import (
	"fmt"
	"io"
	"strings"

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
	for _, t := range cmd.Tracks {
		tags := make([]string, len(t.Tags))

		if len(t.Tags) == 0 {
			continue
		}

		i := 0
		for tag := range t.Tags {
			tags[i] = tag
			i++
		}

		fmt.Fprintf(cmd.Corpus, strings.Join(tags, " "))
		fmt.Fprintf(cmd.Corpus, "\n")
	}

	return nil
}
