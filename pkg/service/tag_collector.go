package service

import "sync"

/*
TagCollector is service that collects all registered tags

*/
type TagCollector struct {
	in  chan []string
	inw *sync.WaitGroup
}

/*
NewTagCollector creates a new instance of TagCollector

*/
func NewTagCollector(wg *sync.WaitGroup) *TagCollector {
	return &TagCollector{
		in:  make(chan []string),
		inw: wg,
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
		print(tag[0] + ": " + tag[1] + "\n")
		c.inw.Done()
	}
}

/*
Wait awaits for all tags to be processed

*/
func (c *TagCollector) Wait() {
	c.inw.Wait()
}
