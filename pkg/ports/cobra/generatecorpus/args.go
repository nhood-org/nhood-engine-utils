package generatecorpus

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type commandArguments struct {
	root                string
	output              string
	thresholdTag        float64
	thresholdSimilarity float64
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

	thresholdTagInt, err := strconv.Atoi(cmd.Flag(thresholdTagsFlagName).Value.String())
	if err != nil {
		return commandArguments{}, errors.New("threshold flag is invalid")
	}
	thresholdTag := float64(thresholdTagInt)

	thresholdSimilarity, err := strconv.ParseFloat(cmd.Flag(thresholdSimilarityFlagName).Value.String(), 64)
	if err != nil {
		return commandArguments{}, errors.New("similar-id-threshold flag is invalid")
	}

	corpusOut := cmd.Flag(outputFlagName).Value.String()

	return commandArguments{
		root:                root,
		output:              corpusOut,
		thresholdTag:        thresholdTag,
		thresholdSimilarity: thresholdSimilarity,
	}, nil
}
