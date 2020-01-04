package gtracks

import (
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/shopspring/decimal"
)

type trackGenerated struct {
	ID       string
	Metadata []decimal.Decimal
}

type trackGenerator interface {
	generate(in model.Track) (trackGenerated, error)
}

type generateVectorsEnvironment struct {
	args      *generateTracksCommandArguments
	generator trackGenerator
}

func (t *trackGenerated) isOnlyZeros() bool {
	zero := decimal.NewFromInt32(0)
	for _, m := range t.Metadata {
		if !zero.Equals(m) {
			return false
		}
	}
	return true
}
