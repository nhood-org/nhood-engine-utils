package command

import (
	"io"

	"github.com/ynqa/wego/pkg/builder"
	"github.com/ynqa/wego/pkg/model/glove"
)

type GenerateGlobalVectorsCmd struct {
	Size        int
	Corpus      io.Reader
	VectorsFile string
}

type GenerateGlobalVectorsHandler interface {
	Handle(GenerateGlobalVectorsCmd) error
}

type generateGlobalVectorsCommandHandler struct {
}

func NewGenerateGlobalVectorsCommandHandler() GenerateGlobalVectorsHandler {
	return generateGlobalVectorsCommandHandler{}
}

func (h generateGlobalVectorsCommandHandler) Handle(cmd GenerateGlobalVectorsCmd) error {
	m, err := builder.NewGloveBuilder().
		Window(5).
		Solver(glove.SGD).
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
