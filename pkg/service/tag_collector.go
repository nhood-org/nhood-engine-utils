package service

import "sync"

/*
TagCollector is service that collects all tags
provided through an input channel 'Input'

*/
type TagCollector struct {
	Input chan []string
	wg    *sync.WaitGroup
}

/*
NewTagCollector creates a new instance of TagCollector

*/
func NewTagCollector(wg *sync.WaitGroup) *TagCollector {
	return &TagCollector{
		Input: make(chan []string),
		wg:    wg,
	}
}

/*
Monitor runs an infinite loop handling incoming tags

*/
func (c *TagCollector) Monitor() error {
	for {
		tag := <-c.Input
		print(tag[0] + ": " + tag[1] + "\n")
		c.wg.Done()
	}
}
