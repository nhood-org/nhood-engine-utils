package main

import (
	"fmt"

	"github.com/nhood-org/nhood-engine-utils/pkg/command"
	"github.com/spf13/cobra"
)

func main() {
	defer handleErrors()
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(command.NewTagCollectorCommand())
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func handleErrors() {
	if r := recover(); r != nil {
		fmt.Println("Error: could not execute command because of an error:", r)
	}
}
