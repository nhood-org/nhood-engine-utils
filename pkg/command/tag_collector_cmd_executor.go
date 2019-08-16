package command

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var workers = newTagCollectorWorkersPool()

func execute(_ *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmdArgs)
	if err != nil {
		panic(err)
	}
	workers.run()
	go handleRootPath(args.Root)
	workers.finalize()
}

type tagCollectorCommandArguments struct {
	Root string
}

func resolveArguments(args []string) (*tagCollectorCommandArguments, error) {
	if len(args) == 0 {
		return nil, errors.New("directory argument is required")
	}

	root := args[0]
	if _, err := os.Stat(root); err == nil {

	} else if os.IsNotExist(err) {
		return nil, errors.New("directory '" + root + "' does not exist")
	} else {
		return nil, errors.New("could not check '" + root + "' directory")
	}

	return &tagCollectorCommandArguments{
		Root: root,
	}, nil
}

func handleRootPath(path string) {
	err := filepath.Walk(path, handlePath)
	if err != nil {
		panic(err)
	}
	workers.done()
}

func handlePath(path string, info os.FileInfo, _ error) error {
	j := tagCollectorWorkerJob{
		path: path,
		info: info,
	}
	workers.addJob(&j)
	return nil
}
