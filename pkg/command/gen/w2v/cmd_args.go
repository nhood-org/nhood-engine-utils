package w2v

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

type word2VecCommandArguments struct {
	Root   string
	Output string
}

func resolveArguments(cmd *cobra.Command, args []string) (*word2VecCommandArguments, error) {
	if len(args) == 0 {
		return nil, errors.New("tracks file argument is required")
	}

	root := args[0]
	if _, err := os.Stat(root); err == nil {

	} else if os.IsNotExist(err) {
		return nil, errors.New("file '" + root + "' does not exist")
	} else {
		return nil, errors.New("could not check '" + root + "' file")
	}

	output := cmd.Flag(outputFlagName).Value.String()

	return &word2VecCommandArguments{
		Root:   root,
		Output: output,
	}, nil
}
