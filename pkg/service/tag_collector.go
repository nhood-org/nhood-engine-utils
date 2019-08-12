package service

import (
	"log"
	"sync"
)

/*
TagCollector is service that collects all registered tags

*/
type TagCollector struct {
	in      chan []string
	inw     *sync.WaitGroup
	process *sync.WaitGroup
}

/*
NewTagCollector creates a new instance of TagCollector

*/
func NewTagCollector() *TagCollector {
	var inw sync.WaitGroup
	var process sync.WaitGroup

	process.Add(1)
	return &TagCollector{
		in:      make(chan []string),
		inw:     &inw,
		process: &process,
	}
}

/*
Register adds incoming tag to the processing channel

*/
func (c *TagCollector) Register(tag []string) error {
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
		log.Println(tag[0] + ": " + tag[1] + "\n")
		c.inw.Done()
	}
}

/*
Done closes input channel

*/
func (c *TagCollector) Done() {
	c.process.Done()
}

/*
Wait awaits for all tags to be processed

*/
func (c *TagCollector) Wait() {
	c.process.Wait()
	c.inw.Wait()
}
