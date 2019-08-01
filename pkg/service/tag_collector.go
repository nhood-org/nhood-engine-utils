package service

/*
TagCollector is service that collects all tags
provided through an input channel 'Input'

*/
type TagCollector struct {
	Input chan []string
}

/*
NewTagCollector creates a new instance of TagCollector

*/
func NewTagCollector() *TagCollector {
	return &TagCollector{
		Input: make(chan []string),
	}
}

/*
Monitor runs an infinite loop handling incoming tags

*/
func (c *TagCollector) Monitor() error {
	for {
		tag := <-c.Input
		print(tag[0] + ": " + tag[1] + "\n")
	}
}
