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
	size          int
	jobs          chan Job
	jobMonitor    *sync.WaitGroup
	jobOKSignal   chan bool
	jobOKCounter  int
	jobNOKCounter int
}

func newWorkerPool(size int) *workerPool {
	var jobMonitor sync.WaitGroup
	return &workerPool{
		size:          size,
		jobs:          make(chan Job),
		jobMonitor:    &jobMonitor,
		jobOKSignal:   make(chan bool),
		jobOKCounter:  0,
		jobNOKCounter: 0,
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
	close(t.jobOKSignal)
}

func (t *workerPool) worker() {
	for j := range t.jobs {
		err := j.Handle()
		if err != nil {
			t.jobOKSignal <- false
		} else {
			t.jobOKSignal <- true
		}
	}
}

func (t *workerPool) monitorJobStatus() {
	for status := range t.jobOKSignal {
		if status {
			t.jobOKCounter++
		} else {
			t.jobNOKCounter++
		}
		if (t.jobOKCounter > 0 && t.jobOKCounter%1000 == 0) || (t.jobNOKCounter > 0 && t.jobNOKCounter%1000 == 0) {
			log.Println("Processed:", t.jobOKCounter, "jobs with [OK] status")
			log.Println("Processed:", t.jobNOKCounter, "jobs with [NOK] status")
		}
		t.jobMonitor.Done()
	}
}
