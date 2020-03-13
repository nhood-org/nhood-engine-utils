package main

import (
	"fmt"
	"log"

	"net/http"
	_ "net/http/pprof"

	"github.com/nhood-org/nhood-engine-utils/pkg/command/w2v"
	"github.com/spf13/cobra"
)

const appName = "engine-utils"

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	defer handleErrors()

	var rootCmd = &cobra.Command{Use: appName}
	rootCmd.AddCommand(w2v.NewCommand())
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
