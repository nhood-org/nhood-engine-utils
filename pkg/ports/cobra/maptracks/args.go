package maptracks

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

const (
	modeFull = "FULL"
	modeTags = "TAGS"
)

type commandArguments struct {
	root   string
	output string
}

func resolveArguments(cmd *cobra.Command, args []string) (commandArguments, error) {
	if len(args) == 0 {
		err := errors.New("tracks directory argument is required")
		panic(err)
	}

	root := args[0]
	if _, err := os.Stat(root); err == nil {
	} else if os.IsNotExist(err) {
		return commandArguments{}, errors.New("file '" + root + "' does not exist")
	} else {
		return commandArguments{}, errors.New("could not check '" + root + "' file")
	}

	output := cmd.Flag(outputFlagName).Value.String()

	return commandArguments{
		root:   root,
		output: output,
	}, nil
}
