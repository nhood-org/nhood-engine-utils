package gvec

import (
	"github.com/nhood-org/nhood-engine-utils/pkg/command/gen/gvec/relevance"
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

type tagFilter interface {
	filter(matrix *relevance.Matrix) *relevance.Matrix
}

type tagRelevanycMapper interface {
	resolve(tags *model.Tags, tracks *model.Tracks) *relevance.Matrix
}

type clusterResolver interface {
	resolve(matrix *relevance.Matrix) [][]string
}

type vectorResolver interface {
	resolve(matrix *relevance.Matrix, clusters [][]string) *model.MetadataVectors
}

type generateVectorsEnvironment struct {
	args               *generateVectorsCommandArguments
	tagRelevanycMapper tagRelevanycMapper
	tagFilter          tagFilter
	clusterResolver    clusterResolver
	vectorResolver     vectorResolver
}
