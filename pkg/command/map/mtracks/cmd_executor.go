package mtracks

import (
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
	"github.com/spf13/cobra"
)

const defaultPoolSize = 1000

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	tags, err := model.ReadTagsFromFile(args.TagsInput)
	if err != nil {
		panic(err)
	}

	env := trackMapperEnvironment{
		args:      args,
		knownTags: tags,
		collector: newTrackCollectorService(),
	}
	env.initialize()

	jobFactory := trackMapperJobFactory{
		env: &env,
	}

	pathWalker := utils.NewPathWalker(defaultPoolSize, args.Root, jobFactory)
	pathWalker.Execute()

	env.finalize()
}
