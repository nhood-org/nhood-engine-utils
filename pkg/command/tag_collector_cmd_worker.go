package command

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/service"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
)

const supportedExtension = ".json"
const defaultPoolSize = 1000

type tagCollectorWorkerJob struct {
	path string
	info os.FileInfo
}

type tagCollectorWorkersPool struct {
	size          int
	env           *tagCollectorWorkersEnvironment
	jobs          chan tagCollectorWorkerJob
	jobsMonitor   *sync.WaitGroup
	jobOKSignal   chan bool
	jobOKCounter  int
	jobNOKCounter int
}

type tagCollectorCommandArguments struct {
	Root           string
	Output         string
	CountThreshold int
}

type tagCollectorWorkersEnvironment struct {
	collector *service.TagCollector
	args      *tagCollectorCommandArguments
}

func newTagCollectorWorkersPool(env *tagCollectorWorkersEnvironment) *tagCollectorWorkersPool {
	var jobsMonitor sync.WaitGroup
	return &tagCollectorWorkersPool{
		size:          defaultPoolSize,
		env:           env,
		jobs:          make(chan tagCollectorWorkerJob),
		jobsMonitor:   &jobsMonitor,
		jobOKSignal:   make(chan bool),
		jobOKCounter:  0,
		jobNOKCounter: 0,
	}
}

func (t *tagCollectorWorkersPool) run() {
	for w := 0; w < t.size; w++ {
		go t.worker()
	}
	go t.env.collector.Monitor()
	go t.monitorJobStatus()
}

func (t *tagCollectorWorkersPool) worker() {
	for j := range t.jobs {
		t.handleJob(&j)
		t.jobsMonitor.Done()
	}
}

func (t *tagCollectorWorkersPool) addJob(job *tagCollectorWorkerJob) {
	t.jobsMonitor.Add(1)
	t.jobs <- *job
}

func (t *tagCollectorWorkersPool) done() {
	t.jobsMonitor.Wait()
	t.env.collector.Close()
}

func (t *tagCollectorWorkersPool) finalize() {
	t.env.collector.Wait()

	tags, err := t.env.collector.GetResults()
	if err != nil {
		panic(err)
	}

	out, err := utils.NewOutputFile(t.env.args.Output)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = out.Close()
		if err != nil {
			panic(err)
		}
	}()

	count := 0
	for _, t := range tags {
		line := fmt.Sprint("Tag:", t.Name,
			"; Count:", t.Statistics.Count(),
			"; Weight:Avg:", t.Statistics.Avg(),
			"; Weight:Max:", t.Statistics.Max(),
			"; Weight:Min:", t.Statistics.Min())
		err = out.Append(line)
		if err != nil {
			panic(err)
		}
		count++
	}

	fmt.Printf("Collected %d tags\n", count)
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

func (t *tagCollectorWorkersPool) handleRootPath(path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		return t.handlePath(path, info, err)
	})
	if err != nil {
		panic(err)
	}
	t.done()
}

func (t *tagCollectorWorkersPool) handlePath(path string, info os.FileInfo, _ error) error {
	j := tagCollectorWorkerJob{
		path: path,
		info: info,
	}
	t.addJob(&j)
	return nil
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
	err = t.env.collector.Register(&tag)
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
