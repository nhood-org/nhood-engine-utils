package w2v

import (
	"github.com/spf13/cobra"
)

const outputFlagName = "output"
const outputDefault = "word2vec.out"

const vectorSizeFlagName = "size"
const vectorSizeDefault = 15

const thresholdSimilarityFlagName = "similarity-threshold"
const thresholdSimilarityDefault = 0.3

const thresholdTagsFlagName = "tag-threshold"
const thresholdTagsDefault = 30

/*
NewCommand returns an instance of a cobra.Command
implementing a track metadata operations

*/
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "w2v [tracks files to process]",
		Short: "Find and generate all the tracks in the directory into tracks with metadata vectors",
		Long:  `w2v is for mapping of all JSON track files in the given directory.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   execute,
	}
	cmd.Flags().StringP(outputFlagName, "o", outputDefault, "output file")
	cmd.Flags().Uint(vectorSizeFlagName, vectorSizeDefault, "generated vector size")
	cmd.Flags().Float64P(thresholdSimilarityFlagName, "s", thresholdSimilarityDefault, "track similarity weight threshold. Similar track IDs with lower weight will be ignored")
	cmd.Flags().UintP(thresholdTagsFlagName, "t", thresholdTagsDefault, "tag weight threshold. Tags with lower weight will be ignored")
	return cmd
}
