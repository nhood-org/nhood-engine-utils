package generatecorpus

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const (
	tracksOutputDefault = "tracks.out"
)

type generateCorpusCommandArguments struct {
	Root                string
	Output              string
	TracksOutput        string
	ThresholdTag        int
	ThresholdSimilarity float64
}

func resolveArguments(cmd *cobra.Command, args []string) (generateCorpusCommandArguments, error) {
	if len(args) == 0 {
		return generateCorpusCommandArguments{}, errors.New("tracks directory argument is required")
	}

	root := args[0]
	if _, err := os.Stat(root); err == nil {

	} else if os.IsNotExist(err) {
		return generateCorpusCommandArguments{}, errors.New("file '" + root + "' does not exist")
	} else {
		return generateCorpusCommandArguments{}, errors.New("could not check '" + root + "' file")
	}

	thresholdTag, err := strconv.Atoi(cmd.Flag("tag-threshold").Value.String())
	if err != nil {
		return generateCorpusCommandArguments{}, errors.New("threshold flag is invalid")
	}

	thresholdSimilarity, err := strconv.ParseFloat(cmd.Flag("similarity-threshold").Value.String(), 64)
	if err != nil {
		return generateCorpusCommandArguments{}, errors.New("similar-id-threshold flag is invalid")
	}

	output := cmd.Flag(outputFlagName).Value.String()

	return generateCorpusCommandArguments{
		Root:                root,
		Output:              output,
		TracksOutput:        tracksOutputDefault,
		ThresholdTag:        thresholdTag,
		ThresholdSimilarity: thresholdSimilarity,
	}, nil
}
