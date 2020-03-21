package generatecorpus

import "github.com/spf13/cobra"

const outputFlagName = "output"
const outputDefault = "corpus.out"

const thresholdSimilarityFlagName = "similarity-threshold"
const thresholdSimilarityDefault = 0.3

const thresholdTagsFlagName = "tag-threshold"
const thresholdTagsDefault = 30

/*
NewCommand returns an instance of a cobra.Command

*/
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-corpus [tracks files to process]",
		Short: "Find all the tracks within the directory and transform into corpus file",
		Long:  `generate-corpus is used for mapping of all JSON track files in the given directory and transforming those into a single corpus file for further usage.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   execute,
	}
	cmd.Flags().StringP(outputFlagName, "o", outputDefault, "output file")
	cmd.Flags().Float64P(thresholdSimilarityFlagName, "s", thresholdSimilarityDefault, "track similarity weight threshold. Similar track IDs with lower weight will be ignored")
	cmd.Flags().UintP(thresholdTagsFlagName, "t", thresholdTagsDefault, "tag weight threshold. Tags with lower weight will be ignored")
	return cmd
}
