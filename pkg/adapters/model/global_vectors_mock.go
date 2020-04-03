package model

import (
	"io"
)

type GlobalVectorsResolverMock struct {
}

func NewGlobalVectorsResolverMock() GlobalVectorsResolverMock {
	return GlobalVectorsResolverMock{}
}

func (g GlobalVectorsResolverMock) Resolve(in io.Reader, out io.Writer) error {
	_, err := io.Copy(out, in)
	return err
}
