package generatecorpus

import (
	"io"
	"os"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/app/command"
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
	"github.com/spf13/cobra"
)

const (
	commandName                 = "generate-corpus"
	outputFlagName              = "output"
	outputDefault               = "corpus.out"
	thresholdSimilarityFlagName = "similarity-threshold"
	thresholdSimilarityDefault  = 0.3
	thresholdTagsFlagName       = "tag-threshold"
	thresholdTagsDefault        = 30
	modeFlagName                = "mode"
	modeFlagDefault             = modeTags
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   commandName + " [tracks files to process]",
		Short: "Find all the tracks within the directory and transform into corpus file",
		Long: commandName +
			" is used for mapping of all JSON track files in the given directory and transforming those into a single corpus file for further usage.",
		Args: cobra.MinimumNArgs(1),
		Run:  execute,
	}
	cmd.Flags().StringP(
		outputFlagName, "o", outputDefault, "output file")
	cmd.Flags().StringP(
		modeFlagName, "m", modeFlagDefault,
		"mode defines what data is included in the output corpus\navailable modes:\n- FULL - all tag data is included\n- TAGS - only tags are included")
	cmd.Flags().Float64P(
		thresholdSimilarityFlagName, "s", thresholdSimilarityDefault,
		"track similarity weight threshold. Similar track IDs with lower weight will be ignored")
	cmd.Flags().UintP(
		thresholdTagsFlagName, "t", thresholdTagsDefault,
		"tag weight threshold. Tags with lower weight will be ignored")
	return cmd
}

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	tracks := readInputTracks(args.root)

	corpus, err := os.Create(args.output)
	if err != nil {
		panic(err)
	}
	defer corpus.Close()

	switch args.mode {
	case modeFull:
		executeFullMode(tracks, corpus, args)
	case modeTags:
		executeTagsMode(tracks, corpus, args)
	}
}

func executeFullMode(
	tracks []model.Track,
	corpus io.Writer,
	args commandArguments,
) {
	c := command.GenerateCorpusCmd{
		Tracks:              tracks,
		ThresholdTag:        args.thresholdTag,
		ThresholdSimilarity: args.thresholdSimilarity,
		Corpus:              corpus,
	}
	err := command.NewGenerateCorpusCommandHandler().Handle(c)
	if err != nil {
		panic(err)
	}
}

func executeTagsMode(
	tracks []model.Track,
	corpus io.Writer,
	args commandArguments,
) {
	c := command.GenerateCorpusTagsCmd{
		Tracks: tracks,
		Corpus: corpus,
	}
	err := command.NewGenerateCorpusTagsCommandHandler().Handle(c)
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
