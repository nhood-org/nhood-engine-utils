package gtracks

import (
	"fmt"
	"log"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
	"github.com/spf13/cobra"
)

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	tracks, err := model.ReadTracksFromFile(args.Root)
	if err != nil {
		panic(err)
	}

	vectors, err := model.ReadMetadataVectorsFromFile(args.VectorsInput)
	if err != nil {
		panic(err)
	}

	generator := newTrackGeneratorImpl(*vectors)

	env := generateVectorsEnvironment{
		args:      args,
		generator: generator,
	}

	i := 0
	generated := make([]trackGenerated, 0)
	for _, t := range *tracks {
		g, err := env.generator.generate(t)
		if err != nil {
			panic(err)
		}
		generated = append(generated, g)
		if i%1000 == 0 {
			log.Println("Mapped:", i, "tracks")
		}
		i++
	}

	filtered := make([]trackGenerated, 0)
	for _, g := range generated {
		if !g.isOnlyZeros() {
			filtered = append(filtered, g)
		}
	}

	utils.SaveToFile(filtered, args.Output)

	ratio := (len(filtered) * 100) / len(generated)
	fmt.Printf("Generated %d tracks, where %d%% are not empty.\n", len(generated), ratio)
}
