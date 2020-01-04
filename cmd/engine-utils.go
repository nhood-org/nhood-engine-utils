package main

import (
	"fmt"
	"log"

	"github.com/nhood-org/nhood-engine-utils/pkg/command/gen/gtracks"
	"github.com/nhood-org/nhood-engine-utils/pkg/command/gen/gvec"
	"github.com/nhood-org/nhood-engine-utils/pkg/command/map/mtags"
	"github.com/nhood-org/nhood-engine-utils/pkg/command/map/mtracks"

	"net/http"
	_ "net/http/pprof"

	"github.com/spf13/cobra"
)

const appName = "engine-utils"

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	defer handleErrors()
	var rootCmd = &cobra.Command{Use: appName}
	rootCmd.AddCommand(gtracks.NewCommand())
	rootCmd.AddCommand(gvec.NewCommand())
	rootCmd.AddCommand(mtags.NewCommand())
	rootCmd.AddCommand(mtracks.NewCommand())
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func handleErrors() {
	if r := recover(); r != nil {
		fmt.Println("Error: could not execute cmd because of an error:", r)
	}
}
