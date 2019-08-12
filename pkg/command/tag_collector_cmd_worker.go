package command

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/service"
)

const supportedExtension = ".json"
const defaultPoolSize = 1000

type tagCollectorWorkerJob struct {
	path string
	info os.FileInfo
}

type tagCollectorWorkersPool struct {
	size       int
	jobs       chan tagCollectorWorkerJob
	jobCounter int
	jobsw      *sync.WaitGroup
	collector  *service.TagCollector
}

func newTagCollectorWorkersPool() *tagCollectorWorkersPool {
	var jobsw sync.WaitGroup
	return &tagCollectorWorkersPool{
		size:       defaultPoolSize,
		jobs:       make(chan tagCollectorWorkerJob),
		jobCounter: 0,
		jobsw:      &jobsw,
		collector:  service.NewTagCollector(),
	}
}

func (t *tagCollectorWorkersPool) run() {
	for w := 0; w < t.size; w++ {
		go t.worker()
	}
	go t.collector.Monitor()
}

func (t *tagCollectorWorkersPool) worker() {
	for j := range t.jobs {
		t.handleJob(&j)
		t.jobsw.Done()
	}
}

func (t *tagCollectorWorkersPool) addJob(job *tagCollectorWorkerJob) {
	t.jobsw.Add(1)
	t.jobs <- *job
}

func (t *tagCollectorWorkersPool) done() {
	t.collector.Done()
}

func (t *tagCollectorWorkersPool) wait() {
	t.collector.Wait()
	t.jobsw.Wait()
}

func (t *tagCollectorWorkersPool) handleJob(job *tagCollectorWorkerJob) error {
	if strings.HasSuffix(job.info.Name(), supportedExtension) {
		err := t.handleJSONPath(job.path)
		if err != nil {
			return err
		}
		t.jobCounter++
		if t.jobCounter%1000 == 0 {
			log.Println("Processed:", t.jobCounter, "files")
		}
	}
	return nil
}

func (t *tagCollectorWorkersPool) handleJSONPath(path string) error {
	track, err := model.ReadTrack(path)
	if err != nil {
		log.Printf("Could not process path: %s because of an error: %v\n", path, err)
		return err
	}
	for _, tag := range track.Tags {
		err := t.collector.Register(tag)
		if err != nil {
			return err
		}
	}
	return nil
}
