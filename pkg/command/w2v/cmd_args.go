package w2v

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type word2VecCommandArguments struct {
	Root                string
	Output              string
	Size                int
	ThresholdTag        int
	ThresholdSimilarity float64
}

func resolveArguments(cmd *cobra.Command, args []string) (word2VecCommandArguments, error) {
	if len(args) == 0 {
		return word2VecCommandArguments{}, errors.New("tracks file argument is required")
	}

	root := args[0]
	if _, err := os.Stat(root); err == nil {

	} else if os.IsNotExist(err) {
		return word2VecCommandArguments{}, errors.New("file '" + root + "' does not exist")
	} else {
		return word2VecCommandArguments{}, errors.New("could not check '" + root + "' file")
	}

	size, err := strconv.Atoi(cmd.Flag(vectorSizeFlagName).Value.String())
	if err != nil {
		return word2VecCommandArguments{}, errors.New("size flag is invalid")
	}

	if size <= 0 {
		return word2VecCommandArguments{}, errors.New("size value must be positive")
	}

	thresholdTag, err := strconv.Atoi(cmd.Flag("tag-threshold").Value.String())
	if err != nil {
		return word2VecCommandArguments{}, errors.New("threshold flag is invalid")
	}

	thresholdSimilarity, err := strconv.ParseFloat(cmd.Flag("similarity-threshold").Value.String(), 64)
	if err != nil {
		return word2VecCommandArguments{}, errors.New("similar-id-threshold flag is invalid")
	}

	output := cmd.Flag(outputFlagName).Value.String()

	return word2VecCommandArguments{
		Root:                root,
		Output:              output,
		Size:                size,
		ThresholdTag:        thresholdTag,
		ThresholdSimilarity: thresholdSimilarity,
	}, nil
}
