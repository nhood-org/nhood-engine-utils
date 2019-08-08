package main

import (
	"github.com/nhood-org/nhood-engine-utils/pkg/command"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(command.NewTagCollectorCommand())
	rootCmd.Execute()
}
