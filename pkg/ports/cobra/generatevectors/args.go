package generatevectors

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type commandArguments struct {
	tracks         string
	tags           string
	tagsVectorSize int
	output         string
}

func resolveArguments(cmd *cobra.Command, args []string) (commandArguments, error) {
	if len(args) == 0 {
		return commandArguments{}, errors.New("corpus file argument is required")
	}

	tracks := args[0]
	if _, err := os.Stat(tracks); err == nil {

	} else if os.IsNotExist(err) {
		return commandArguments{}, errors.New("file '" + tracks + "' does not exist")
	} else {
		return commandArguments{}, errors.New("could not check '" + tracks + "' file")
	}

	tags := args[1]
	if _, err := os.Stat(tags); err == nil {

	} else if os.IsNotExist(err) {
		return commandArguments{}, errors.New("file '" + tags + "' does not exist")
	} else {
		return commandArguments{}, errors.New("could not check '" + tags + "' file")
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
		tracks:         tracks,
		tags:           tags,
		tagsVectorSize: size,
		output:         output,
	}, nil
}
