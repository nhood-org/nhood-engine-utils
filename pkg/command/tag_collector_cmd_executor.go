package command

import (
	"errors"
	"os"
	"strconv"

	"github.com/nhood-org/nhood-engine-utils/pkg/service"
	"github.com/spf13/cobra"
)

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	collectorConfig := service.NewTagCollectorConfig(args.CountThreshold)
	collector := service.NewTagCollector(collectorConfig)
	env := &tagCollectorWorkersEnvironment{
		collector: collector,
		args:      args,
	}
	workers := newTagCollectorWorkersPool(env)

	workers.run()
	go workers.handleRootPath(args.Root)
	workers.finalize()
}

func resolveArguments(cmd *cobra.Command, args []string) (*tagCollectorCommandArguments, error) {
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

	threshold, err := strconv.Atoi(cmd.Flag("threshold").Value.String())
	if err != nil {
		return nil, errors.New("threshold flag is invalid")
	}

	output := cmd.Flag("output").Value.String()

	return &tagCollectorCommandArguments{
		Root:           root,
		Output:         output,
		CountThreshold: threshold,
	}, nil
}
