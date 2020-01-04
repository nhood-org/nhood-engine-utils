package mtracks

import (
	"fmt"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
)

type trackMapperEnvironment struct {
	args      *trackMapperCommandArguments
	knownTags *model.Tags
	collector *trackCollectorService
}

func (t *trackMapperEnvironment) initialize() {
	t.collector.run()
}

func (t *trackMapperEnvironment) finalize() {
	tracks := t.collector.tracks

	err := utils.SaveToFile(tracks, t.args.Output)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Collected %d tracks\n", len(*tracks))
}
