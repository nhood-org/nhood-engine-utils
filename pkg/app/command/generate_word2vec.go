package command

import (
	"io"

	"github.com/pkg/errors"
)

type Word2VecVectorsCommandResolver interface {
	Resolve(size int, in io.Reader, out io.Writer) error
}

type GenerateWord2VecVectorsCmd struct {
	Size   int
	Corpus io.Reader
	Output io.Writer
}

type GenerateWord2VecVectorsHandler interface {
	Handle(GenerateWord2VecVectorsCmd) error
}

type generateWord2VecVectorsCommandHandler struct {
	resolver Word2VecVectorsCommandResolver
}

func NewGenerateWord2VecVectorsCommandHandler(
	resolver Word2VecVectorsCommandResolver,
) GenerateWord2VecVectorsHandler {
	return generateWord2VecVectorsCommandHandler{
		resolver: resolver,
	}
}

func (h generateWord2VecVectorsCommandHandler) Handle(cmd GenerateWord2VecVectorsCmd) error {
	err := h.resolver.Resolve(cmd.Size, cmd.Corpus, cmd.Output)
	if err != nil {
		return errors.Wrap(err, "could not generate word2vec vectors")
	}

	return nil
}
