package gtracks

import (
	"github.com/spf13/cobra"
)

const vectorsInputFlagName = "vectors"
const outputFlagName = "output"

const vectorsInputDefault = "vectors.out.json"
const outputDefault = "tracks-generated.out.json"

/*
NewCommand returns an instance of a cobra.Command
implementing a track metadata operations

*/
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gtracks [tracks file to process]",
		Short: "Find and generate all the tracks in the directory into tracks with metadata vectors",
		Long:  `gtracks is for mapping of all JSON track files in the given directory.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   execute,
	}
	cmd.Flags().StringP(vectorsInputFlagName, "v", vectorsInputDefault, "vectors input file")
	cmd.Flags().StringP(outputFlagName, "o", outputDefault, "output file")
	return cmd
}
