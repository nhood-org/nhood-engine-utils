package command

import (
	"log"

	"github.com/spf13/cobra"
)

/*
NewTagCollectorCommand return an instance of a cobra.Command
implementing a tags collection operations

*/
func NewTagCollectorCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "tags [directory to walk]",
		Short: "Find all the tags in the directory",
		Long:  `tags is for resolution of all tags within the JSON song files in the given directory.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, cmdArgs []string) {
			err := execute(cmd, cmdArgs)
			if err != nil {
				log.Fatalln("Could not resolve tags because of an error:", err)
			}
		},
	}
}
