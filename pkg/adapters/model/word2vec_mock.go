package model

import (
	"io"
)

type Word2VecVectorsResolverMock struct {
}

func NewWord2VecVectorsResolverMock() Word2VecVectorsResolverMock {
	return Word2VecVectorsResolverMock{}
}

func (g Word2VecVectorsResolverMock) Resolve(size int, in io.Reader, out io.Writer) error {
	_, err := io.Copy(out, in)
	return err
}
