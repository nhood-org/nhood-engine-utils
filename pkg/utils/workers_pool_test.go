package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

const testJobCount = 1000
const testWorkerPoolSize = 10

type TestJob struct {
	out chan bool
}

func (t TestJob) Handle() error {
	t.out <- true
	return nil
}

func TestWorkerHandlesAllRegisteredJobs(t *testing.T) {
	a := assert.New(t)

	defer goleak.VerifyNoLeaks(t)

	jobFeedbackChannel := make(chan bool)
	defer close(jobFeedbackChannel)

	jobCount := testJobCount

	workerPoolSize := testWorkerPoolSize
	workers := newWorkerPool(workerPoolSize)
	workers.run()

	count := 0
	go func() {
		for range jobFeedbackChannel {
			count++
		}
	}()

	for i := 0; i < jobCount; i++ {
		j := TestJob{out: jobFeedbackChannel}
		workers.addJob(j)
	}
	workers.done()

	a.Equal(jobCount, count)
}

type TestOKJob struct{}

func (t TestOKJob) Handle() error {
	return nil
}

func TestWorkerCountsOKJobs(t *testing.T) {
	a := assert.New(t)

	defer goleak.VerifyNoLeaks(t)

	jobCount := testJobCount

	workerPoolSize := testWorkerPoolSize
	workers := newWorkerPool(workerPoolSize)
	workers.run()

	for i := 0; i < jobCount; i++ {
		workers.addJob(TestOKJob{})
	}
	workers.done()

	expectedNOKCount := 0
	expectedOKCount := jobCount

	a.Equal(expectedOKCount, workers.jobOKCounter)
	a.Equal(expectedNOKCount, workers.jobNOKCounter)
}

type TestNOKJob struct{}

func (t TestNOKJob) Handle() error {
	return errors.New("job failed")
}

func TestWorkerCountsNOKJobs(t *testing.T) {
	a := assert.New(t)

	defer goleak.VerifyNoLeaks(t)

	jobCount := testJobCount
	jobFailureRatio := 5

	workerPoolSize := testWorkerPoolSize
	workers := newWorkerPool(workerPoolSize)
	workers.run()

	for i := 0; i < jobCount; i++ {
		if i%jobFailureRatio == 0 {
			workers.addJob(TestNOKJob{})
		} else {
			workers.addJob(TestOKJob{})
		}
	}
	workers.done()

	expectedNOKCount := jobCount / jobFailureRatio
	expectedOKCount := jobCount - expectedNOKCount

	a.Equal(expectedOKCount, workers.jobOKCounter)
	a.Equal(expectedNOKCount, workers.jobNOKCounter)
}
