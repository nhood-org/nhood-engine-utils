package generatecorpus

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/nhood-org/nhood-engine-utils/pkg/app/command"
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
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
		Use:   commandName + " [tracks file to process]",
		Short: "Read all the tracks from input file and transform into corpus file",
		Long: commandName +
			" reads all the tracks from input file and transform into corpus file for further usage.",
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

	tracks := readInputTracks(args.input)

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

func readInputTracks(input string) []model.Track {
	inputFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	inputBytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}

	var tracks []model.Track
	err = json.Unmarshal(inputBytes, &tracks)
	if err != nil {
		panic(err)
	}

	return tracks
}
