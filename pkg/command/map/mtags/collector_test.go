package mtags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTagResolverMock struct{}

func (t testTagResolverMock) Resolve(tag string) ([]string, error) {
	return []string{tag}, nil
}

func TestNewTagCollector(t *testing.T) {
	a := assert.New(t)

	countThreshold := 0
	resolver := testTagResolverMock{}
	config := newTagCollectorConfigService(countThreshold, resolver)
	c := newTagCollectorService(config)

	a.NotNil(c)
	a.Equal(config, c.config)

	a.NotNil(c.in)
	a.NotNil(c.inw)
	a.NotNil(c.tags)
}

func TestNewTagCollectorConfig(t *testing.T) {
	a := assert.New(t)

	countThreshold := 10
	resolver := testTagResolverMock{}
	config := newTagCollectorConfigService(countThreshold, resolver)

	a.NotNil(config)
	a.Equal(countThreshold, config.countThreshold)
}

func TestTagCollectorComputesWeightAverages(t *testing.T) {
	a := assert.New(t)

	countThreshold := 0
	resolver := testTagResolverMock{}
	config := newTagCollectorConfigService(countThreshold, resolver)
	c := newTagCollectorService(config)
	c.run()

	err := c.register(&tag{Name: "rock", Weight: 10})
	a.Nil(err)

	err = c.register(&tag{Name: "rock", Weight: 20})
	a.Nil(err)

	err = c.register(&tag{Name: "rock", Weight: 30})
	a.Nil(err)

	c.wait()

	results, err := c.getResults()
	a.Nil(err)

	a.Len(results, 1)
	a.Equal("rock", results[0].Name)
	a.Equal(float64(20), results[0].Statistics.Avg())
	a.Equal(3, results[0].Statistics.Count())
}

func TestTagCollectorReturnsOnlyTagsWithCountAboveThreshold(t *testing.T) {
	a := assert.New(t)

	countThreshold := 2
	resolver := testTagResolverMock{}
	config := newTagCollectorConfigService(countThreshold, resolver)
	c := newTagCollectorService(config)
	c.run()

	err := c.register(&tag{Name: "rock", Weight: 10})
	a.Nil(err)

	err = c.register(&tag{Name: "rock", Weight: 10})
	a.Nil(err)

	err = c.register(&tag{Name: "metal", Weight: 10})
	a.Nil(err)

	c.wait()

	results, err := c.getResults()
	a.Nil(err)

	a.Len(results, 1)
	a.Equal("rock", results[0].Name)
	a.Equal(float64(10), results[0].Statistics.Avg())
	a.Equal(2, results[0].Statistics.Count())
}
