package gtracks

import (
	"fmt"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/shopspring/decimal"
)

// TODO!!! add tests
const maxVectorSize = 32

type trackGeneratorImpl struct {
	vectors model.MetadataVectors
}

func newTrackGeneratorImpl(vectors model.MetadataVectors) *trackGeneratorImpl {
	validateVectors(vectors)
	return &trackGeneratorImpl{
		vectors: vectors,
	}
}

func validateVectors(vectors model.MetadataVectors) error {
	for i, v := range vectors.Clusters {
		if len(v) >= maxVectorSize {
			return fmt.Errorf("size of vector[%d] exceeds or equals max value: %d", i, maxVectorSize)
		}
	}
	return nil
}

func (t *trackGeneratorImpl) generate(in model.Track) (trackGenerated, error) {
	metadataSize := len(t.vectors.Clusters)
	metadata := make([]int32, metadataSize)

	for _, tag := range in.Tags {
		c, ok := t.vectors.TagCoordinates[tag]
		if ok {
			clusterSize := t.vectors.ClusterSizes[c.Cluster]
			maxValue := 1<<clusterSize - 1
			metadata[c.Cluster] = int32(maxValue >> c.ClusterIndex)
		}
	}

	g := trackGenerated{
		ID:       in.ID,
		Metadata: make([]decimal.Decimal, metadataSize),
	}

	for i, m := range metadata {
		g.Metadata[i] = decimal.NewFromInt32(m)
	}

	return g, nil
}
