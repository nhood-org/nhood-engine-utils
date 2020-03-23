package command

import (
	"io"

	"github.com/ynqa/wego/pkg/builder"
	"github.com/ynqa/wego/pkg/model/word2vec"
)

type GenerateWord2VecVectorsCmd struct {
	Size        int
	Corpus      io.Reader
	VectorsFile string
}

type GenerateWord2VecVectorsHandler interface {
	Handle(GenerateWord2VecVectorsCmd) error
}

type generateWord2VecVectorsCommandHandler struct {
}

func NewGenerateWord2VecVectorsCommandHandler() GenerateWord2VecVectorsHandler {
	return generateWord2VecVectorsCommandHandler{}
}

func (h generateWord2VecVectorsCommandHandler) Handle(cmd GenerateWord2VecVectorsCmd) error {
	m, err := builder.NewWord2vecBuilder().
		Dimension(cmd.Size).
		Window(5).
		Model(word2vec.CBOW).
		Optimizer(word2vec.NEGATIVE_SAMPLING).
		NegativeSampleSize(5).
		Verbose().
		Build()
	if err != nil {
		panic(err)
	}

	err = m.Train(cmd.Corpus)
	if err != nil {
		panic(err)
	}

	err = m.Save(cmd.VectorsFile)
	if err != nil {
		panic(err)
	}

	return nil
}
