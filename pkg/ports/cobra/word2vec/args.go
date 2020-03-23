package word2vec

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type commandArguments struct {
	corpus string
	output string
	size   int
}

func resolveArguments(cmd *cobra.Command, args []string) (commandArguments, error) {
	if len(args) == 0 {
		return commandArguments{}, errors.New("corpus file argument is required")
	}

	corpus := args[0]
	if _, err := os.Stat(corpus); err == nil {

	} else if os.IsNotExist(err) {
		return commandArguments{}, errors.New("file '" + corpus + "' does not exist")
	} else {
		return commandArguments{}, errors.New("could not check '" + corpus + "' file")
	}

	size, err := strconv.Atoi(cmd.Flag(vectorSizeFlagName).Value.String())
	if err != nil {
		return commandArguments{}, errors.New("size flag is invalid")
	}

	if size <= 0 {
		return commandArguments{}, errors.New("size value must be positive")
	}

	output := cmd.Flag(outputFlagName).Value.String()

	return commandArguments{
		corpus: corpus,
		output: output,
		size:   size,
	}, nil
}
