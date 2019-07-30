package main

import (
	"fmt"
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
			fmt.Printf("%+v\n", *t)
		case <-time.After(timeout):
			done = true
		}

		if done {
			break
		}
	}
}
