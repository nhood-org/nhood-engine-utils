package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

const testDataDirectory = "./test_data"
const testDataJSONExtension = ".json"
const timeout = 200 * time.Millisecond

var ch = make(chan *model.Track)
var chw sync.WaitGroup

func main() {
	chw.Add(1)
	go handleRootPath(testDataDirectory)
	go monitor()
	chw.Wait()
}

func handleRootPath(path string) {
	filepath.Walk(path, handlePath)
	chw.Done()
}

func handlePath(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(info.Name(), testDataJSONExtension) {
		chw.Add(1)
		handleJSONPath(path)
	}
	return nil
}

func handleJSONPath(path string) {
	t, _ := model.ReadTrack(path)
	ch <- t
	chw.Done()
}

func monitor() {
	for {
		fmt.Printf("%+v\n", *<-ch)
	}
}
