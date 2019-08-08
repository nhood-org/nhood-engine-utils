package command

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/service"

	"github.com/spf13/cobra"
)

const supportedExtension = ".json"

var collector = service.NewTagCollector()

/*
NewTagCollectorCommand return an instance of a cobra.Command
implementing a tags collection operations

*/
func NewTagCollectorCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "tags [directory to walk]",
		Short: "Find all the tags in the directory",
		Long:  `tags is for resolution of all tags within the JSON song files in the given directory.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, cmdArgs []string) {
			err := execute(cmd, cmdArgs)
			if err != nil {
				fmt.Println("Could not resolve tags because of an error:", err)
			}
		},
	}
}

var fileCounter int32

type tagCollectorCommandArguments struct {
	Root string
}

func execute(cmd *cobra.Command, cmdArgs []string) error {
	args, err := resolveArguments(cmdArgs)
	if err != nil {
		return err
	}

	go handleRootPath(args.Root)
	go collector.Monitor()
	collector.Wait()

	return nil
}

func resolveArguments(args []string) (*tagCollectorCommandArguments, error) {
	if len(args) == 0 {
		return nil, errors.New("Directory argument is required")
	}

	root := args[0]
	if _, err := os.Stat(root); err == nil {

	} else if os.IsNotExist(err) {
		return nil, errors.New("Directory '" + root + "' does not exist")
	} else {
		return nil, errors.New("Could not check '" + root + "' directory")
	}

	return &tagCollectorCommandArguments{
		Root: root,
	}, nil
}

func handleRootPath(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		dirPath := path + "/" + f.Name()
		wg.Add(1)
		go handleDirectoryPath(dirPath, &wg)
	}
	wg.Wait()

	collector.Done()
}

func handleDirectoryPath(path string, wg *sync.WaitGroup) {
	filepath.Walk(path, handlePath)
	wg.Done()
}

func handlePath(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(info.Name(), supportedExtension) {
		err := handleJSONPath(path)
		if err != nil {
			return err
		}
		fileCounter++
		if fileCounter%1000 == 0 {
			println("Processed:", fileCounter, "files")
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
