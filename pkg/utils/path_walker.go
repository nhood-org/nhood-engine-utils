package utils

import (
	"os"
	"path/filepath"
)

/*
JobFactory is a simple interface for generation of PathWalker jobs

*/
type JobFactory interface {
	Create(path string, info os.FileInfo) (Job, error)
}

/*
PathWalker walks a file directory and calls a job handler on each of the files detected

*/
type PathWalker struct {
	dir        string
	jobFactory JobFactory
	workers    *workerPool
}

/*
NewPathWalker creates a new instance of PathWalker

*/
func NewPathWalker(size int, dir string, jobFactory JobFactory) *PathWalker {
	workers := newWorkerPool(size)
	return &PathWalker{
		dir:        dir,
		jobFactory: jobFactory,
		workers:    workers,
	}
}

/*
Execute runs scan of directory path

*/
func (p *PathWalker) Execute() {
	p.workers.run()
	err := filepath.Walk(p.dir, func(path string, info os.FileInfo, err error) error {
		j, err := p.jobFactory.Create(path, info)
		if err != nil {
			return err
		}
		p.workers.addJob(j)
		return nil
	})
	if err != nil {
		panic(err)
	}
	p.workers.done()
}
