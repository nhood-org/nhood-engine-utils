package mtags

import (
	"fmt"
	"sort"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/tags"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
)

type tagCollectorService struct {
	config *tagCollectorServiceConfig
	in     chan *tag
	inw    *sync.WaitGroup
	tags   map[string]*utils.MovingStatistics
}

type tagCollectorServiceConfig struct {
	countThreshold int
	resolver       tags.Resolver
}

func newTagCollectorService(config *tagCollectorServiceConfig) *tagCollectorService {
	var inw sync.WaitGroup
	return &tagCollectorService{
		config: config,
		in:     make(chan *tag),
		inw:    &inw,
		tags:   make(map[string]*utils.MovingStatistics),
	}
}

func newTagCollectorConfigService(countThreshold int, resolver tags.Resolver) *tagCollectorServiceConfig {
	return &tagCollectorServiceConfig{
		countThreshold: countThreshold,
		resolver:       resolver,
	}
}

func (c *tagCollectorService) register(tag *tag) error {
	c.inw.Add(1)
	c.in <- tag
	return nil
}

func (c *tagCollectorService) run() {
	go func() {
		for tag := range c.in {
			tags, err := c.config.resolver.Resolve(tag.Name)
			if err != nil {
				fmt.Println("Could not resolve tags from ", tag.Name)
			}

			for _, t := range tags {
				c.handleTag(t, tag.Weight)
			}

			c.inw.Done()
		}
	}()
}

func (c *tagCollectorService) handleTag(name string, weight int64) {
	_, ok := c.tags[name]
	if !ok {
		c.tags[name] = utils.NewMovingStatistics()
	}
	ma := c.tags[name]
	ma.Add(float64(weight))
}

func (c *tagCollectorService) wait() {
	c.inw.Wait()
}

func (c *tagCollectorService) getResults() ([]tagStatistics, error) {
	var tagSlice []tagStatistics
	for k, v := range c.tags {
		if v.Count() >= c.config.countThreshold {
			tagSlice = append(tagSlice, tagStatistics{k, v})
		}
	}

	sort.Slice(tagSlice, func(i, j int) bool {
		return tagSlice[i].Statistics.Count() > tagSlice[j].Statistics.Count()
	})

	return tagSlice, nil
}
