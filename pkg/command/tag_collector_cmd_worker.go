package command

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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
	size          int
	jobs          chan tagCollectorWorkerJob
	jobsw         *sync.WaitGroup
	jobOKSignal   chan bool
	jobOKCounter  int
	jobNOKCounter int
	collector     *service.TagCollector
}

func newTagCollectorWorkersPool() *tagCollectorWorkersPool {
	var jobsw sync.WaitGroup
	return &tagCollectorWorkersPool{
		size:          defaultPoolSize,
		jobs:          make(chan tagCollectorWorkerJob),
		jobsw:         &jobsw,
		jobOKSignal:   make(chan bool),
		jobOKCounter:  0,
		jobNOKCounter: 0,
		collector:     service.NewTagCollector(),
	}
}

func (t *tagCollectorWorkersPool) run() {
	for w := 0; w < t.size; w++ {
		go t.worker()
	}
	go t.collector.Monitor()
	go t.monitorJobStatus()
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
	t.collector.Close()
}

func (t *tagCollectorWorkersPool) finalize() {
	t.collector.Wait()
	t.jobsw.Wait()
	t.collector.PrintResults()
}

func (t *tagCollectorWorkersPool) handleJob(job *tagCollectorWorkerJob) {
	if !strings.HasSuffix(job.info.Name(), supportedExtension) {
		return
	}

	err := t.handleJSONPath(job.path)
	if err != nil {
		log.Printf("Could not process path: %s because of an error: %v\n", job.path, err)
		t.jobOKSignal <- false
	} else {
		t.jobOKSignal <- true
	}
}

func (t *tagCollectorWorkersPool) handleJSONPath(path string) error {
	track, err := model.ReadTrack(path)
	if err != nil {
		return err
	}

	for _, array := range track.Tags {
		err = t.handleTagArray(array)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *tagCollectorWorkersPool) handleTagArray(array []string) error {
	if len(array) != 2 {
		message := fmt.Sprintf("Tag array has illegal size: %d expected: 2\n", len(array))
		return errors.New(message)
	}

	i, err := strconv.ParseInt(array[1], 10, 16)
	if err != nil {
		message := fmt.Sprintf("Could not parse tag weight because of an error: %v\n", err)
		return errors.New(message)
	}

	tag := model.TrackTag{
		Name:   array[0],
		Weight: i,
	}
	err = t.collector.Register(&tag)
	if err != nil {
		return err
	}

	return nil
}

func (t *tagCollectorWorkersPool) monitorJobStatus() {
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
	}
}
