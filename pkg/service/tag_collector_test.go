package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTagCollector(t *testing.T) {
	a := assert.New(t)

	config := &TagCollectorConfig{}
	c := NewTagCollector(config)

	a.NotNil(c)
	a.Equal(c.config, config)

	a.NotNil(c.in)
	a.NotNil(c.inw)
	a.NotNil(c.closeSignal)
	a.False(c.closed)
	a.NotNil(c.tags)
}

func TestNewTagCollectorConfig(t *testing.T) {
	a := assert.New(t)

	countThreshold := 10
	config := NewTagCollectorConfig(countThreshold)

	a.NotNil(config)
	a.Equal(config.countThreshold, countThreshold)
}

func TestTagCollectorDoesNotAcceptTagsWhenClosed(t *testing.T) {

}

func TestTagCollectorDoesNotAcceptTagsWithCommonWordName(t *testing.T) {

}

func TestTagCollectorDoesNotAcceptTagsWithSingleDigitName(t *testing.T) {

}

func TestTagCollectorDoesNotAcceptTagsWithSingleCharacterName(t *testing.T) {

}

func TestTagCollectorIsCaseInsensitive(t *testing.T) {

}

func TestTagCollectorAcceptsSingleWordTagNames(t *testing.T) {

}

func TestTagCollectorAcceptsMultiWordTagNames(t *testing.T) {

}

func TestTagCollectorComputesWeightAverages(t *testing.T) {

}
