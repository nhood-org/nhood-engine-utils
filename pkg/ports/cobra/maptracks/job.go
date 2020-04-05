package maptracks

import (
	"os"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
)

const defaultPoolSize = 8

type job struct {
	out   chan model.RawTrack
	outWg *sync.WaitGroup
	path  string
	info  os.FileInfo
}

func (j job) Handle() error {
	if j.info.IsDir() {
		return nil
	}

	raw, err := model.ReadRawTrackFromFile(j.path)
	if err != nil {
		return err
	}

	j.outWg.Add(1)
	j.out <- *raw

	return nil
}

type jobFactory struct {
	out   chan model.RawTrack
	outWg *sync.WaitGroup
}

func (f jobFactory) Create(path string, info os.FileInfo) (utils.Job, error) {
	return job{
		out:   f.out,
		outWg: f.outWg,
		path:  path,
		info:  info,
	}, nil
}
