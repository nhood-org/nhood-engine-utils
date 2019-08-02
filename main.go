package main

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/arguments"
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/service"
)

const supportedExtension = ".json"

var mwg sync.WaitGroup
var cwg sync.WaitGroup

var collector = service.NewTagCollector(&cwg)

func main() {
	args, err := arguments.ResolveArguments()
	if err != nil {
		panic(err)
	}

	mwg.Add(1)
	go handleRootPath(args.Root)
	go collector.Monitor()

	mwg.Wait()
	collector.Wait()
}

func handleRootPath(path string) {
	filepath.Walk(path, handlePath)
	mwg.Done()
}

func handlePath(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(info.Name(), supportedExtension) {
		err := handleJSONPath(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleJSONPath(path string) error {
	t, _ := model.ReadTrack(path)
	for _, tag := range t.Tags {
		err := collector.Register(tag)
		if err != nil {
			return err
		}
	}
	return nil
}
