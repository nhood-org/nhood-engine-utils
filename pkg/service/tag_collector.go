package service

import (
	"errors"
	"sort"
	"strings"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
)

/*
TagCollector is a service that collects all the registered tags and processes them

*/
type TagCollector struct {
	config      *TagCollectorConfig
	in          chan *model.TrackTag
	inw         *sync.WaitGroup
	closeSignal chan bool
	closed      bool
	tags        map[string]*utils.MovingAverage
}

/*
TagCollectorConfig defines an internal TagCollectorConfig configuration

*/
type TagCollectorConfig struct {
	countThreshold int
}

/*
Tag defines a tag and its statistics

*/
type Tag struct {
	Name       string
	Statistics *utils.MovingAverage
}

/*
NewTagCollector creates a new instance of a TagCollector

*/
func NewTagCollector(config *TagCollectorConfig) *TagCollector {
	var inw sync.WaitGroup
	return &TagCollector{
		config:      config,
		in:          make(chan *model.TrackTag),
		inw:         &inw,
		closeSignal: make(chan bool),
		closed:      false,
		tags:        make(map[string]*utils.MovingAverage),
	}
}

/*
NewTagCollectorConfig creates a new instance of a TagCollectorConfig

*/
func NewTagCollectorConfig(countThreshold int) *TagCollectorConfig {
	return &TagCollectorConfig{
		countThreshold: countThreshold,
	}
}

/*
Register adds incoming tag to the processing channel
After an input channel is closed tags will not be accepted

*/
func (c *TagCollector) Register(tag *model.TrackTag) error {
	if c.closed {
		return errors.New("input channel is already closed")
	}

	c.inw.Add(1)
	c.in <- tag

	return nil
}

/*
Monitor runs an infinite loop handling incoming tags

*/
func (c *TagCollector) Monitor() {
	for {
		select {
		case tag := <-c.in:
			tags := strings.Split(tag.Name, " ")
			for _, t := range tags {
				name := strings.ToLower(t)
				isValid := true
				isValid = isValid && nameIsNotAuxiliaryWord(name)
				isValid = isValid && nameIsNotASingleCharacter(name)
				if isValid {
					c.handleTag(name, tag.Weight)
				}
			}
			c.inw.Done()
		case _ = <-c.closeSignal:
			c.closeSignal <- true
			break
		}
	}
}

func (c *TagCollector) handleTag(name string, weight int64) {
	_, ok := c.tags[name]
	if !ok {
		c.tags[name] = utils.NewMovingAverage()
	}
	ma := c.tags[name]
	ma.Add(float64(weight))
}

/*
Close closes input channel execution
After an input channel is closed tags may not be processed
Channel cannot be open again

*/
func (c *TagCollector) Close() {
	c.closeSignal <- true
}

/*
Wait awaits for a close signal and then for
all registered tags to be processed
l
*/
func (c *TagCollector) Wait() {
	c.closed = <-c.closeSignal
	c.inw.Wait()
}

/*
GetResults returns all collected tags with its statistics

*/
func (c *TagCollector) GetResults() ([]Tag, error) {
	if !c.closed {
		return nil, errors.New("input channel is not closed yet")
	}

	var tagSlice []Tag
	for k, v := range c.tags {
		if v.Count() >= c.config.countThreshold {
			tagSlice = append(tagSlice, Tag{k, v})
		}
	}

	sort.Slice(tagSlice, func(i, j int) bool {
		return tagSlice[i].Statistics.Count() > tagSlice[j].Statistics.Count()
	})

	return tagSlice, nil
}
