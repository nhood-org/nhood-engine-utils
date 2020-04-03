package word2vec

import (
	"os"

	"github.com/nhood-org/nhood-engine-utils/pkg/adapters/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/app/command"
	"github.com/spf13/cobra"
)

const (
	commandName        = "word2vec"
	outputFlagName     = "output"
	outputDefault      = "vectors.out"
	vectorSizeFlagName = "size"
	vectorSizeDefault  = 15
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   commandName + " [corpus file to process]",
		Short: "Generate metadata vectors out from the input corpus file",
		Long:  commandName + ` will interpret a whole input corpus file and generate vectors representing each of the words found.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   execute,
	}
	cmd.Flags().StringP(outputFlagName, "o", outputDefault, "output file")
	cmd.Flags().Uint(vectorSizeFlagName, vectorSizeDefault, "generated vector size")
	return cmd
}

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	corpus, err := os.Open(args.corpus)
	if err != nil {
		panic(err)
	}
	defer corpus.Close()

	output, err := os.Create(args.output)
	if err != nil {
		panic(err)
	}
	defer corpus.Close()

	c := command.GenerateWord2VecVectorsCmd{
		Size:   args.size,
		Corpus: corpus,
		Output: output,
	}

	resolver := model.NewWord2VecVectorsResolver()
	err = command.NewGenerateWord2VecVectorsCommandHandler(resolver).Handle(c)
	if err != nil {
		panic(err)
	}
}
