package model

import (
	"io"

	"github.com/pkg/errors"
	"github.com/ynqa/wego/pkg/builder"
	"github.com/ynqa/wego/pkg/model/glove"
)

type GlobalVectorsResolver struct {
	builder *builder.GloveBuilder
}

func NewGlobalVectorsResolver() GlobalVectorsResolver {
	b := builder.NewGloveBuilder().
		Window(5).
		Solver(glove.SGD).
		Verbose()
	return GlobalVectorsResolver{
		builder: b,
	}
}

func (g GlobalVectorsResolver) Resolve(in io.Reader, out io.Writer) error {
	m, err := g.builder.Build()
	if err != nil {
		return errors.Wrap(err, "could not build global vectors model")
	}

	err = m.Train(in)
	if err != nil {
		return errors.Wrap(err, "could not train global vectors model with corpus data")
	}

	err = m.Save(out)
	if err != nil {
		return errors.Wrap(err, "could not save resolved vectors to the output writer")
	}

	return nil
}
