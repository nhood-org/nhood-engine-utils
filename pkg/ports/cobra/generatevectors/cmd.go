package generatevectors

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/nhood-org/nhood-engine-utils/pkg/app/command"
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/spf13/cobra"
)

const (
	commandName        = "generate-vectors"
	outputFlagName     = "output"
	outputDefault      = "vectors.out.csv"
	vectorSizeFlagName = "size"
	vectorSizeDefault  = 10
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   commandName + " [tracks file to process] [tags vectors file]",
		Short: "Generate metadata vectors out from the input tracks file and tags vectors",
		Long:  commandName + ` will traverse a whole input tracks file and generate output with vectors representing each of the words found.`,
		Args:  cobra.MinimumNArgs(2),
		Run:   execute,
	}
	cmd.Flags().StringP(outputFlagName, "o", outputDefault, "output file")
	cmd.Flags().Uint(vectorSizeFlagName, vectorSizeDefault, "tag vector size")
	return cmd
}

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	tracks := readInputTracks(args.tracks)
	tags := readInputTags(args.tags, args.tagsVectorSize)

	output, err := os.Create(args.output)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	c := command.ComputeTrackVectorsCmd{
		Tracks:     tracks,
		Tags:       tags,
		VectorSize: args.tagsVectorSize,
		Output:     output,
	}

	err = command.NewComputeTrackVectorsCommandHandler().Handle(c)
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

func readInputTags(input string, vectorSize int) []model.Tag {
	inputFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	tags := make([]model.Tag, 0)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		tags = append(tags, resolveInputTag(scanner.Text(), vectorSize))
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	return tags
}

func resolveInputTag(input string, vectorSize int) model.Tag {
	inputStrings := strings.Split(input, " ")
	if len(inputStrings) < vectorSize+1 {
		message := fmt.Sprintf(
			"invalid input data: expected tag data '%s' with vector of size at least %d", input, vectorSize)
		panic(message)
	}

	tag := inputStrings[0]

	vector := make([]float64, vectorSize)
	for i := range vector {
		sIdx := i + 1
		s := inputStrings[sIdx]
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			message := fmt.Sprintf(
				"invalid input data of tag %s: could not parse value at index %d as float: %s: %s", tag, sIdx, s, err.Error())
			panic(message)
		}
		vector[i] = f
	}

	return model.Tag{
		Name:   tag,
		Vector: vector,
	}
}
