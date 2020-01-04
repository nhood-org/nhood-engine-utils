package mtags

import (
	"github.com/spf13/cobra"
)

/*
NewCommand returns an instance of a cobra.Command
implementing a tags collection operations

*/
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mtags [directory to walk]",
		Short: "Find and map all the tags in the directory",
		Long:  `mtags is for resolution of all tags within the JSON song files in the given directory.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   execute,
	}
	cmd.Flags().UintP("threshold", "t", 3000, "count threshold. Tags with lower count will not be collected")
	cmd.Flags().StringP("output", "o", "tags.out.json", "output file")
	return cmd
}
