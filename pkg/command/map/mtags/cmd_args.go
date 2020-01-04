package mtags

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type tagCollectorCommandArguments struct {
	Root           string
	Output         string
	CountThreshold int
}

func resolveArguments(cmd *cobra.Command, args []string) (*tagCollectorCommandArguments, error) {
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

	threshold, err := strconv.Atoi(cmd.Flag("threshold").Value.String())
	if err != nil {
		return nil, errors.New("threshold flag is invalid")
	}

	output := cmd.Flag("output").Value.String()

	return &tagCollectorCommandArguments{
		Root:           root,
		Output:         output,
		CountThreshold: threshold,
	}, nil
}
