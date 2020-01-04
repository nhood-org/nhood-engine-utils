package gtracks

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

type generateTracksCommandArguments struct {
	Root         string
	Output       string
	VectorsInput string
}

func resolveArguments(cmd *cobra.Command, args []string) (*generateTracksCommandArguments, error) {
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

	vectorsInput := cmd.Flag(vectorsInputFlagName).Value.String()
	if _, err := os.Stat(vectorsInput); err == nil {

	} else if os.IsNotExist(err) {
		return nil, errors.New("file '" + vectorsInput + "' does not exist")
	} else {
		return nil, errors.New("could not check '" + vectorsInput + "' input file")
	}

	output := cmd.Flag(outputFlagName).Value.String()

	return &generateTracksCommandArguments{
		Root:         root,
		Output:       output,
		VectorsInput: vectorsInput,
	}, nil
}
