package generatecorpus

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const (
	modeFull = "FULL"
	modeTags = "TAGS"
)

type commandArguments struct {
	input               string
	output              string
	mode                string
	thresholdTag        float64
	thresholdSimilarity float64
}

func resolveArguments(cmd *cobra.Command, args []string) (commandArguments, error) {
	if len(args) == 0 {
		err := errors.New("tracks directory argument is required")
		panic(err)
	}

	input := args[0]
	if _, err := os.Stat(input); err == nil {
	} else if os.IsNotExist(err) {
		return commandArguments{}, errors.New("file '" + input + "' does not exist")
	} else {
		return commandArguments{}, errors.New("could not check '" + input + "' file")
	}

	var mode string

	modeVal := cmd.Flag(modeFlagName).Value.String()
	switch modeVal {
	case modeFull:
		mode = modeFull
	case modeTags:
		mode = modeTags
	default:
		return commandArguments{}, errors.New("unknown mode: " + modeVal)
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
		input:               input,
		output:              corpusOut,
		mode:                mode,
		thresholdTag:        thresholdTag,
		thresholdSimilarity: thresholdSimilarity,
	}, nil
}
