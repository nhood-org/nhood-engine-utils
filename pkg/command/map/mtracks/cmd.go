package mtracks

import (
	"github.com/spf13/cobra"
)

/*
NewCommand returns an instance of a cobra.Command
implementing a track mapping operations

*/
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mtracks [directory to walk]",
		Short: "Find and map all the tracks in the directory",
		Long:  `mtracks is for mapping of all JSON track files in the given directory.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   execute,
	}
	cmd.Flags().Float64P("similarity-threshold", "s", 0.3, "track similarity weight threshold. Similar track IDs with lower weight will be ignored")
	cmd.Flags().UintP("tag-threshold", "t", 30, "tag weight threshold. Tags with lower weight will be ignored")
	cmd.Flags().StringP("tags", "i", "tags.out.json", "tag input file")
	cmd.Flags().StringP("output", "o", "tracks.out.json", "output file")
	return cmd
}
