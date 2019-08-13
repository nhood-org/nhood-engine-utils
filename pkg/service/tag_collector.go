package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
)

/*
TagCollector is a service that collects all the registered tags and processes them

*/
type TagCollector struct {
	in          chan *model.TrackTag
	inw         *sync.WaitGroup
	closeSignal chan bool
	closed      bool
	tags        map[string]*utils.MovingAverage
}

/*
NewTagCollector creates a new instance of a TagCollector

*/
func NewTagCollector() *TagCollector {
	var inw sync.WaitGroup
	return &TagCollector{
		in:          make(chan *model.TrackTag),
		inw:         &inw,
		closeSignal: make(chan bool),
		closed:      false,
		tags:        make(map[string]*utils.MovingAverage),
	}
}

/*
Register adds incoming tag to the processing channel
After an input channel is closed tags will not be accepted

*/
func (c *TagCollector) Register(tag *model.TrackTag) error {
	if c.closed {
		return errors.New("Input channel is already closed")
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
		c.handleTag(tag)
		c.inw.Done()
	}
}

func (c *TagCollector) handleTag(tag *model.TrackTag) {
	_, ok := c.tags[tag.Name]
	if !ok {
		c.tags[tag.Name] = utils.NewMovingAverage()
	}
	ma := c.tags[tag.Name]
	ma.Add(float64(tag.Weight))
}

/*
Close closes input channel execution
After an input channel is closed tags may not be processed
Cannel cannot be open again

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
	for key, value := range c.tags {
		if value.Count() < 1000 {
			continue
		}
		fmt.Println("Tag:", key, "; Count:", value.Count(), "; Weight:", value.Avg())
	}
}
