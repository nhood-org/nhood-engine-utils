package utils

import (
	"log"
	"sync"
)

/*
Job is a simple interface for PathWalker jobs

*/
type Job interface {
	Handle() error
}

type workerPool struct {
	size            int
	jobs            chan Job
	jobMonitor      *sync.WaitGroup
	jobReturnSignal chan error
	jobAllCounter   int
	jobErrorCounter int
}

func newWorkerPool(size int) *workerPool {
	var jobMonitor sync.WaitGroup
	return &workerPool{
		size:            size,
		jobs:            make(chan Job),
		jobMonitor:      &jobMonitor,
		jobReturnSignal: make(chan error),
		jobAllCounter:   0,
		jobErrorCounter: 0,
	}
}

func (t *workerPool) run() {
	for w := 0; w < t.size; w++ {
		go t.worker()
	}
	go t.monitorJobStatus()
}

func (t *workerPool) addJob(job Job) {
	t.jobMonitor.Add(1)
	t.jobs <- job
}

func (t *workerPool) done() {
	t.jobMonitor.Wait()
	close(t.jobs)
	close(t.jobReturnSignal)
}

func (t *workerPool) worker() {
	for j := range t.jobs {
		t.jobReturnSignal <- j.Handle()
	}
}

func (t *workerPool) monitorJobStatus() {
	for err := range t.jobReturnSignal {
		t.jobAllCounter++
		if err != nil {
			t.jobErrorCounter++
			log.Println("error: " + err.Error())
		}
		if t.jobAllCounter > 0 && t.jobAllCounter%1000 == 0 {
			log.Println("Processed:", t.jobAllCounter, "jobs")
			log.Println("Processed:", t.jobErrorCounter, "jobs with error")
		}
		t.jobMonitor.Done()
	}
}
