package word2vec

import (
	"github.com/spf13/cobra"
)

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	generateVectors(args)
}
