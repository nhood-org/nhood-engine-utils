package service

import (
	"errors"
	"testing"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/stretchr/testify/assert"
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
	a := assert.New(t)

	config := &TagCollectorConfig{}
	c := NewTagCollector(config)
	go c.Monitor()

	err := c.Register(&model.TrackTag{})
	a.Nil(err)

	c.Close()
	c.Wait()

	err = c.Register(&model.TrackTag{})
	a.Equal(err, errors.New("input channel is already closed"))
}

func TestTagCollectorDoesNotReturnResultsWhenNotClosed(t *testing.T) {
	a := assert.New(t)

	config := &TagCollectorConfig{}
	c := NewTagCollector(config)

	results, err := c.GetResults()
	a.Nil(results)
	a.Equal(err, errors.New("input channel is not closed yet"))
}

func TestTagCollectorDoesNotAcceptTagsWithCommonWordName(t *testing.T) {
	a := assert.New(t)

	config := &TagCollectorConfig{}
	c := NewTagCollector(config)
	go c.Monitor()

	err := c.Register(&model.TrackTag{Name: "a", Weight: 100})
	a.Nil(err)

	err = c.Register(&model.TrackTag{Name: "the", Weight: 100})
	a.Nil(err)

	err = c.Register(&model.TrackTag{Name: "I", Weight: 100})
	a.Nil(err)

	c.Close()
	c.Wait()

	results, err := c.GetResults()
	a.Nil(err)
	a.Empty(results)
}

func TestTagCollectorDoesNotAcceptTagsWithSingleDigitName(t *testing.T) {
	a := assert.New(t)

	config := &TagCollectorConfig{}
	c := NewTagCollector(config)
	go c.Monitor()

	err := c.Register(&model.TrackTag{Name: "1", Weight: 100})
	a.Nil(err)

	err = c.Register(&model.TrackTag{Name: "0", Weight: 100})
	a.Nil(err)

	c.Close()
	c.Wait()

	results, err := c.GetResults()
	a.Nil(err)
	a.Empty(results)
}

func TestTagCollectorDoesNotAcceptTagsWithSingleCharacterName(t *testing.T) {
	a := assert.New(t)

	config := &TagCollectorConfig{}
	c := NewTagCollector(config)
	go c.Monitor()

	err := c.Register(&model.TrackTag{Name: "a", Weight: 100})
	a.Nil(err)

	err = c.Register(&model.TrackTag{Name: "z", Weight: 100})
	a.Nil(err)

	c.Close()
	c.Wait()

	results, err := c.GetResults()
	a.Nil(err)
	a.Empty(results)
}

func TestTagCollectorIsCaseInsensitive(t *testing.T) {
	a := assert.New(t)

	config := &TagCollectorConfig{}
	c := NewTagCollector(config)
	go c.Monitor()

	err := c.Register(&model.TrackTag{Name: "rock", Weight: 100})
	a.Nil(err)

	err = c.Register(&model.TrackTag{Name: "Rock", Weight: 100})
	a.Nil(err)

	err = c.Register(&model.TrackTag{Name: "rOcK", Weight: 100})
	a.Nil(err)

	c.Close()
	c.Wait()

	results, err := c.GetResults()
	a.Nil(err)

	a.Len(results, 1)
	a.Equal(results[0].Name, "rock")
}

func TestTagCollectorAcceptsMultiWordTagNames(t *testing.T) {
	a := assert.New(t)

	config := &TagCollectorConfig{}
	c := NewTagCollector(config)
	go c.Monitor()

	err := c.Register(&model.TrackTag{Name: "american rock", Weight: 100})
	a.Nil(err)

	c.Close()
	c.Wait()

	results, err := c.GetResults()
	a.Nil(err)

	a.Len(results, 2)

	names := make([]string, 0)
	for _, t := range results {
		names = append(names, t.Name)
	}
	a.Contains(names, "rock")
	a.Contains(names, "american")
}

func TestTagCollectorComputesWeightAverages(t *testing.T) {
	a := assert.New(t)

	config := &TagCollectorConfig{}
	c := NewTagCollector(config)
	go c.Monitor()

	err := c.Register(&model.TrackTag{Name: "rock", Weight: 10})
	a.Nil(err)

	err = c.Register(&model.TrackTag{Name: "rock", Weight: 20})
	a.Nil(err)

	err = c.Register(&model.TrackTag{Name: "rock", Weight: 30})
	a.Nil(err)

	c.Close()
	c.Wait()

	results, err := c.GetResults()
	a.Nil(err)

	a.Len(results, 1)
	a.Equal(results[0].Name, "rock")
	a.Equal(results[0].Statistics.Avg(), float64(20))
	a.Equal(results[0].Statistics.Count(), 3)
}
