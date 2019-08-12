package command

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var workers = newTagCollectorWorkersPool()

func execute(cmd *cobra.Command, cmdArgs []string) error {
	args, err := resolveArguments(cmdArgs)
	if err != nil {
		return err
	}
	workers.run()
	go handleRootPath(args.Root)
	workers.wait()
	return nil
}

type tagCollectorCommandArguments struct {
	Root string
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
	filepath.Walk(path, handlePath)
	workers.done()
}

func handlePath(path string, info os.FileInfo, err error) error {
	j := tagCollectorWorkerJob{
		path: path,
		info: info,
	}
	workers.addJob(&j)
	return nil
}
