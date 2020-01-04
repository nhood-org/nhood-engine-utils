package tags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitterIsCaseInsensitive(t *testing.T) {
	a := assert.New(t)

	s := NewSplitter()

	for _, in := range []string{"rock", "Rock", "rOcK"} {
		t.Run(in, func(t *testing.T) {
			results, err := s.Resolve(in)
			a.Nil(err)
			a.Len(results, 1)
			a.Contains(results, "rock")
		})
	}
}

func TestSplitterAcceptsMultiWordTagNames(t *testing.T) {
	a := assert.New(t)

	s := NewSplitter()

	results, err := s.Resolve("american rock")
	a.Nil(err)
	a.Len(results, 2)
	a.Contains(results, "american")
	a.Contains(results, "rock")
}

func TestSplitterDoesNotAcceptTagsWithCommonWordName(t *testing.T) {
	a := assert.New(t)

	s := NewSplitter()

	for in := range auxiliaryWords {
		t.Run(in, func(t *testing.T) {
			results, err := s.Resolve(in)
			a.Nil(err)
			a.Empty(results)
		})
	}
}

func TestSplitterDoesNotAcceptTagsWithSingleDigitName(t *testing.T) {
	a := assert.New(t)

	s := NewSplitter()

	for _, in := range []string{"1", "0"} {
		t.Run(in, func(t *testing.T) {
			results, err := s.Resolve(in)
			a.Nil(err)
			a.Empty(results)
		})
	}
}

func TestSplitterDoesNotAcceptTagsWithSingleCharacterName(t *testing.T) {
	a := assert.New(t)

	s := NewSplitter()

	for _, in := range []string{"a", "b", "y", "z"} {
		t.Run(in, func(t *testing.T) {
			results, err := s.Resolve(in)
			a.Nil(err)
			a.Empty(results)
		})
	}
}
