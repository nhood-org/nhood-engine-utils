package service

import (
	"errors"
	"fmt"
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

type tag struct {
	name string
	ma   *utils.MovingAverage
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
		tag := <-c.in
		tags := strings.Split(tag.Name, " ")
		for _, t := range tags {
			name := strings.ToLower(t)
			isValid := nameIsNotAuxiliaryWord(name)
			if isValid {
				c.handleTag(name, tag.Weight)
			}
		}
		c.inw.Done()
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

*/
func (c *TagCollector) Wait() {
	_ = <-c.closeSignal
	c.inw.Wait()
}

/*
PrintResults prints all collected tags with its average weights

*/
func (c *TagCollector) PrintResults() {
	sorted := c.sortTagsByCount(c.tags)
	printTags(sorted)
}

func (c *TagCollector) sortTagsByCount(tags map[string]*utils.MovingAverage) []tag {
	var tagSlice []tag
	for k, v := range tags {
		if v.Count() >= float64(c.config.countThreshold) {
			tagSlice = append(tagSlice, tag{k, v})
		}
	}

	sort.Slice(tagSlice, func(i, j int) bool {
		return tagSlice[i].ma.Count() > tagSlice[j].ma.Count()
	})

	return tagSlice
}

func printTags(tagSlice []tag) {
	for _, t := range tagSlice {
		fmt.Println("Tag:", t.name, "; Count:", t.ma.Count(), "; Weight:", t.ma.Avg())
	}
}
