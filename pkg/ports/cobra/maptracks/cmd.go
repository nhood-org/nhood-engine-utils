package maptracks

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
	"github.com/spf13/cobra"
)

const (
	commandName    = "map-tracks"
	outputFlagName = "output"
	outputDefault  = "tracks.out"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   commandName + " [tracks files to process]",
		Short: "Find all the tracks within the directory and transform into one single file",
		Long: commandName +
			" is used for mapping of all JSON track files in the given directory and transforming those into a single file for further usage.",
		Args: cobra.MinimumNArgs(1),
		Run:  execute,
	}
	cmd.Flags().StringP(outputFlagName, "o", outputDefault, "tracks output file")
	return cmd
}

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	tracks := readInputTracks(args.root)
	tracksBytes, err := json.Marshal(tracks)
	if err != nil {
		panic(err)
	}

	output, err := os.Create(args.output)
	if err != nil {
		panic(err)
	}

	defer output.Close()

	_, err = output.Write(tracksBytes)
	if err != nil {
		panic(err)
	}
}

func readInputTracks(root string) []model.Track {
	var outWg sync.WaitGroup
	out := make(chan model.RawTrack)

	jobFactory := jobFactory{
		outWg: &outWg,
		out:   out,
	}

	tracks := make([]model.Track, 0)
	go func() {
		for {
			o := <-out
			track := model.TrackFromRaw(o)
			tracks = append(tracks, track)
			outWg.Done()
		}
	}()

	pathWalker := utils.NewPathWalker(defaultPoolSize, root, jobFactory)
	pathWalker.Execute()

	outWg.Wait()

	return tracks
}
