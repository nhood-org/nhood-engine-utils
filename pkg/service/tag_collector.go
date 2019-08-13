package service

import (
	"errors"
	"log"
	"sync"
)

/*
TagCollector is a service that collects all the registered tags and processes them

*/
type TagCollector struct {
	in          chan []string
	inw         *sync.WaitGroup
	closeSignal chan bool
	closed      bool
}

/*
NewTagCollector creates a new instance of a TagCollector

*/
func NewTagCollector() *TagCollector {
	var inw sync.WaitGroup
	return &TagCollector{
		in:          make(chan []string),
		inw:         &inw,
		closeSignal: make(chan bool),
		closed:      false,
	}
}

/*
Register adds incoming tag to the processing channel
After an input channel is closed tags will not be accepted

*/
func (c *TagCollector) Register(tag []string) error {
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

func (c *TagCollector) handleTag(tag []string) {
	log.Println(tag[0] + ": " + tag[1] + "\n")
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
