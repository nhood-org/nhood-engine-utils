package mtags

import (
	"github.com/nhood-org/nhood-engine-utils/pkg/tags"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
	"github.com/spf13/cobra"
)

const defaultPoolSize = 1000

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	tagsResolver := tags.NewSanitizer()
	collectorConfig := newTagCollectorConfigService(args.CountThreshold, tagsResolver)
	collector := newTagCollectorService(collectorConfig)
	env := tagCollectorEnvironment{
		collector: collector,
		args:      args,
	}

	jobFactory := tagCollectionJobFactory{
		env: &env,
	}

	env.initialize()

	pathWalker := utils.NewPathWalker(defaultPoolSize, args.Root, jobFactory)
	pathWalker.Execute()

	env.finalize()
}
