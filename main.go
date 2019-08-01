package main

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/nhood-org/nhood-engine-utils/pkg/arguments"
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/service"
)

const testDataJSONExtension = ".json"
const timeout = 200 * time.Millisecond

var wg sync.WaitGroup
var collector = service.NewTagCollector(&wg)

func main() {
	args, err := arguments.ResolveArguments()
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go handleRootPath(args.Root)
	go collector.Monitor()
	wg.Wait()
}

func handleRootPath(path string) {
	filepath.Walk(path, handlePath)
	wg.Done()
}

func handlePath(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(info.Name(), testDataJSONExtension) {
		handleJSONPath(path)
	}
	return nil
}

func handleJSONPath(path string) {
	t, _ := model.ReadTrack(path)
	for _, tag := range t.Tags {
		wg.Add(1)
		collector.Input <- tag
	}
}
