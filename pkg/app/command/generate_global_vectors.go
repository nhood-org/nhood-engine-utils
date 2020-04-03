package command

import (
	"io"

	"github.com/pkg/errors"
)

type GlobalVectorsResolver interface {
	Resolve(in io.Reader, out io.Writer) error
}

type GenerateGlobalVectorsCmd struct {
	Corpus io.Reader
	Output io.Writer
}

type GenerateGlobalVectorsCommandHandler interface {
	Handle(GenerateGlobalVectorsCmd) error
}

type generateGlobalVectorsCommandHandler struct {
	resolver GlobalVectorsResolver
}

func NewGenerateGlobalVectorsCommandHandler(
	resolver GlobalVectorsResolver,
) GenerateGlobalVectorsCommandHandler {
	return generateGlobalVectorsCommandHandler{
		resolver: resolver,
	}
}

func (h generateGlobalVectorsCommandHandler) Handle(cmd GenerateGlobalVectorsCmd) error {
	err := h.resolver.Resolve(cmd.Corpus, cmd.Output)
	if err != nil {
		return errors.Wrap(err, "could not generate global vectors")
	}

	return nil
}
