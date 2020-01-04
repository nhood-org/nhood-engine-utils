package gvec

import (
	"log"
	"sort"

	"github.com/nhood-org/nhood-engine-utils/pkg/command/gen/gvec/relevance"
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

type vectorResolverImpl struct {
}

func newVectorResolverImpl() *vectorResolverImpl {
	return &vectorResolverImpl{}
}

func (v *vectorResolverImpl) resolve(matrix *relevance.Matrix, clusters [][]string) *model.MetadataVectors {
	sortedClusters := make([][]string, 0)
	for i, c := range clusters {
		sorted := sortCluster(matrix, c)
		sortedClusters = append(sortedClusters, sorted)
		log.Println("Cluster:", i, "mapped into vector")
	}
	return model.NewMetadataVectors(sortedClusters)
}

func sortCluster(matrix *relevance.Matrix, cluster []string) []string {
	sort.Slice(cluster, func(i, j int) bool {
		ti := cluster[i]
		si := matrix.Weight(ti)
		tj := cluster[j]
		sj := matrix.Weight(tj)
		return si > sj
	})

	clusterR := make([]string, 0)
	clusterL := make([]string, 0)

	for i := range cluster {
		if i%2 == 0 {
			clusterR = append(clusterR, cluster[i])
		} else {
			clusterL = append(clusterL, cluster[i])
		}
	}

	sorted := clusterR
	for i := len(clusterL) - 1; i >= 0; i-- {
		sorted = append(sorted, clusterL[i])
	}

	return sorted
}
