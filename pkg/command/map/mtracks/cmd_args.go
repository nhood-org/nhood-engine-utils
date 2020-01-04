package mtracks

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type trackMapperCommandArguments struct {
	Root                string
	Output              string
	TagsInput           string
	SimilarityThreshold float64
	TagThreshold        int
}

func resolveArguments(cmd *cobra.Command, args []string) (*trackMapperCommandArguments, error) {
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

	similarityThreshold, err := strconv.ParseFloat(cmd.Flag("similarity-threshold").Value.String(), 64)
	if err != nil {
		return nil, errors.New("similar-id-threshold flag is invalid")
	}

	tagThreshold, err := strconv.Atoi(cmd.Flag("tag-threshold").Value.String())
	if err != nil {
		return nil, errors.New("threshold flag is invalid")
	}

	tagsInput := cmd.Flag("tags").Value.String()
	if _, err := os.Stat(tagsInput); err == nil {

	} else if os.IsNotExist(err) {
		return nil, errors.New("file '" + tagsInput + "' does not exist")
	} else {
		return nil, errors.New("could not check '" + tagsInput + "' input file")
	}

	output := cmd.Flag("output").Value.String()

	return &trackMapperCommandArguments{
		Root:                root,
		Output:              output,
		TagsInput:           tagsInput,
		SimilarityThreshold: similarityThreshold,
		TagThreshold:        tagThreshold,
	}, nil
}
