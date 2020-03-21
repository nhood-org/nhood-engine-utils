package word2vec

import (
	"github.com/spf13/cobra"
)

const outputFlagName = "output"
const outputDefault = "vectors.out"

const vectorSizeFlagName = "size"
const vectorSizeDefault = 15

const thresholdSimilarityFlagName = "similarity-threshold"
const thresholdSimilarityDefault = 0.3

const thresholdTagsFlagName = "tag-threshold"
const thresholdTagsDefault = 30

/*
NewCommand returns an instance of a cobra.Command

*/
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "word2vec [corpus file to process]",
		Short: "Generate metadata vectors out from the input corpus file",
		Long:  `word2vec will interpret a whole input corpus file and generate vectors representing each of the words found.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   execute,
	}
	cmd.Flags().StringP(outputFlagName, "o", outputDefault, "output file")
	cmd.Flags().Uint(vectorSizeFlagName, vectorSizeDefault, "generated vector size")
	return cmd
}
