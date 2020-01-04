package gvec

import "github.com/nhood-org/nhood-engine-utils/pkg/command/gen/gvec/relevance"

type tagFilterImplConfig struct {
	minimalTagWeight int
}

type tagFilterImpl struct {
	config *tagFilterImplConfig
}

func newTagFilterImplConfig(minimalTagWeight int) *tagFilterImplConfig {
	return &tagFilterImplConfig{
		minimalTagWeight: minimalTagWeight,
	}
}

func newTagFilterImpl(config *tagFilterImplConfig) *tagFilterImpl {
	return &tagFilterImpl{
		config: config,
	}
}

func (f *tagFilterImpl) filter(matrix *relevance.Matrix) *relevance.Matrix {
	filteredTags := make([]string, 0)
	for _, t := range matrix.Tags() {
		if matrix.Weight(t) > f.config.minimalTagWeight {
			filteredTags = append(filteredTags, t)
		}
	}
	matrix.SetTags(filteredTags)
	return matrix
}
