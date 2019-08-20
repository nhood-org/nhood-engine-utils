package command

import (
	"github.com/spf13/cobra"
)

/*
NewTagCollectorCommand return an instance of a cobra.Command
implementing a tags collection operations

*/
func NewTagCollectorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tags [directory to walk]",
		Short: "Find all the tags in the directory",
		Long:  `tags is for resolution of all tags within the JSON song files in the given directory.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   execute,
	}
	cmd.Flags().UintP("threshold", "t", 3000, "count threshold. Tags with lower count will not be displayed")

	return cmd
}
