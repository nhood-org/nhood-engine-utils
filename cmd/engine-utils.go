package main

import (
	"fmt"
	"log"

	"net/http"
	_ "net/http/pprof"

	"github.com/nhood-org/nhood-engine-utils/pkg/ports/cobra/generatecorpus"
	"github.com/nhood-org/nhood-engine-utils/pkg/ports/cobra/glove"
	"github.com/nhood-org/nhood-engine-utils/pkg/ports/cobra/maptracks"
	"github.com/nhood-org/nhood-engine-utils/pkg/ports/cobra/word2vec"
	"github.com/spf13/cobra"
)

const appName = "engine-utils"

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	defer handleErrors()

	var rootCmd = &cobra.Command{Use: appName}
	rootCmd.AddCommand(generatecorpus.NewCommand())
	rootCmd.AddCommand(glove.NewCommand())
	rootCmd.AddCommand(maptracks.NewCommand())
	rootCmd.AddCommand(word2vec.NewCommand())
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
