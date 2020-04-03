package model

import (
	"io"

	"github.com/pkg/errors"
	"github.com/ynqa/wego/pkg/builder"
	"github.com/ynqa/wego/pkg/model/word2vec"
)

type Word2VecVectorsResolver struct {
	builder *builder.Word2vecBuilder
}

func NewWord2VecVectorsResolver() Word2VecVectorsResolver {
	b := builder.NewWord2vecBuilder().
		Window(5).
		Model(word2vec.CBOW).
		Optimizer(word2vec.NEGATIVE_SAMPLING).
		NegativeSampleSize(5).
		Verbose()
	return Word2VecVectorsResolver{
		builder: b,
	}
}

func (g Word2VecVectorsResolver) Resolve(size int, in io.Reader, out io.Writer) error {
	m, err := g.builder.Dimension(size).Build()
	if err != nil {
		return errors.Wrap(err, "could not build word2vec model")
	}

	err = m.Train(in)
	if err != nil {
		return errors.Wrap(err, "could not train word2vec model with corpus data")
	}

	err = m.Save(out)
	if err != nil {
		return errors.Wrap(err, "could not save resolved vectors to the output writer")
	}

	return nil
}
