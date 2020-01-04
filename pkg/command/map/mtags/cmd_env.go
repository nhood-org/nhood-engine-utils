package mtags

import (
	"fmt"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
)

type tagCollectorEnvironment struct {
	collector *tagCollectorService
	args      *tagCollectorCommandArguments
}

func (t *tagCollectorEnvironment) initialize() {
	t.collector.run()
}

func (t *tagCollectorEnvironment) finalize() {
	t.collector.wait()

	tags, err := t.collector.getResults()
	if err != nil {
		panic(err)
	}

	out := model.Tags{}
	for _, t := range tags {
		out[t.Name] = true
	}

	err = utils.SaveToFile(out, t.args.Output)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Collected %d tags\n", len(out))
}
