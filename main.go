package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

const testDataDirectory = "./test-data"
const testDataJSONExtension = ".json"
const timeout = 200 * time.Millisecond

var ch = make(chan *model.Track)

func main() {
	filepath.Walk(testDataDirectory, handlePath)
	monitor()
}

func handlePath(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(info.Name(), testDataJSONExtension) {
		// TODO!!! use channel in order to not lose files
		go handleJSONPath(path)
	}
	return nil
}

func handleJSONPath(path string) {
	t, _ := model.ReadTrack(path)
	ch <- t
}

func monitor() {
	var done = false

	for {
		select {
		case t := <-ch:
			println(t.ToString())
		case <-time.After(timeout):
			done = true
		}

		if done {
			break
		}
	}
}
