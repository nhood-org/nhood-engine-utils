package mtracks

import (
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

type trackCollectorService struct {
	in     chan *model.Track
	inw    *sync.WaitGroup
	tracks *model.Tracks
}

func newTrackCollectorService() *trackCollectorService {
	var inw sync.WaitGroup
	return &trackCollectorService{
		in:     make(chan *model.Track),
		inw:    &inw,
		tracks: &model.Tracks{},
	}
}

func (c *trackCollectorService) register(track *model.Track) error {
	c.inw.Add(1)
	c.in <- track
	return nil
}

func (c *trackCollectorService) run() {
	go func() {
		for track := range c.in {
			(*c.tracks)[track.ID] = *track
			c.inw.Done()
		}
	}()
}

func (c *trackCollectorService) wait() {
	c.inw.Wait()
}
