package word2vec

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type word2VecCommandArguments struct {
	Corpus string
	Output string
	Size   int
}

func resolveArguments(cmd *cobra.Command, args []string) (word2VecCommandArguments, error) {
	if len(args) == 0 {
		return word2VecCommandArguments{}, errors.New("corpus file argument is required")
	}

	corpus := args[0]
	if _, err := os.Stat(corpus); err == nil {

	} else if os.IsNotExist(err) {
		return word2VecCommandArguments{}, errors.New("file '" + corpus + "' does not exist")
	} else {
		return word2VecCommandArguments{}, errors.New("could not check '" + corpus + "' file")
	}

	size, err := strconv.Atoi(cmd.Flag(vectorSizeFlagName).Value.String())
	if err != nil {
		return word2VecCommandArguments{}, errors.New("size flag is invalid")
	}

	if size <= 0 {
		return word2VecCommandArguments{}, errors.New("size value must be positive")
	}

	output := cmd.Flag(outputFlagName).Value.String()

	return word2VecCommandArguments{
		Corpus: corpus,
		Output: output,
		Size:   size,
	}, nil
}
