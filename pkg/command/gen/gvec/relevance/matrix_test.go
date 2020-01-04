package relevance

import (
	"testing"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMatrixIsProperlyCreated(t *testing.T) {
	a := assert.New(t)

	tags := &model.Tags{
		"T1": true,
		"T2": true,
		"T3": true,
	}
	m := NewMatrix(tags)
	a.NotNil(m)
	a.Equal(3, m.Len())

	a.Equal(0, m.Relevance("T1", "T1"))
	a.Equal(0, m.Relevance("T1", "T2"))
	a.Equal(0, m.Relevance("T1", "T3"))
	a.Equal(0, m.Weight("T1"))

	a.Equal(0, m.Relevance("T2", "T1"))
	a.Equal(0, m.Relevance("T2", "T2"))
	a.Equal(0, m.Relevance("T2", "T3"))
	a.Equal(0, m.Weight("T2"))

	a.Equal(0, m.Relevance("T3", "T1"))
	a.Equal(0, m.Relevance("T3", "T2"))
	a.Equal(0, m.Relevance("T3", "T3"))
	a.Equal(0, m.Weight("T3"))

	a.Equal(0, m.MaxTagsCount())
}

func TestRelevanceMatrixIsIncremented(t *testing.T) {
	a := assert.New(t)

	tags := &model.Tags{
		"T1": true,
		"T2": true,
		"T3": true,
	}
	m := NewMatrix(tags)
	a.Equal(3, m.Len())

	m.Increment("T1", "T2")
	m.Increment("T1", "T2")
	m.Increment("T2", "T3")

	a.Equal(0, m.Relevance("T1", "T1"))
	a.Equal(2, m.Relevance("T1", "T2"))
	a.Equal(0, m.Relevance("T1", "T3"))
	a.Equal(2, m.Weight("T1"))

	a.Equal(0, m.Relevance("T2", "T1"))
	a.Equal(0, m.Relevance("T2", "T2"))
	a.Equal(1, m.Relevance("T2", "T3"))
	a.Equal(1, m.Weight("T2"))

	a.Equal(2, m.MaxTagsCount())
}
