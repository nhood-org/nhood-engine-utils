package model

import (
	"encoding/json"
	"os"
)

/*
MetadataTagCoordinate defines coordinates
to the tag location within the clusters

*/
type MetadataTagCoordinate struct {
	Cluster      int
	ClusterIndex int
}

/*
MetadataVectors defines metadata vectors entity

*/
type MetadataVectors struct {
	TagCoordinates map[string]MetadataTagCoordinate
	Clusters       [][]string
	ClusterSizes   []int
}

/*
NewMetadataVectors return new instance of MetadataVectors

*/
func NewMetadataVectors(clusters [][]string) *MetadataVectors {
	v := &MetadataVectors{
		TagCoordinates: make(map[string]MetadataTagCoordinate),
		Clusters:       clusters,
		ClusterSizes:   make([]int, len(clusters)),
	}
	v.resolveClusterCoordinates()
	v.resolveClusterSizes()
	return v
}

func (m *MetadataVectors) resolveClusterCoordinates() {
	for i := 0; i < len(m.Clusters); i++ {
		cluster := m.Clusters[i]
		for j := 0; j < len(cluster); j++ {
			t := m.Clusters[i][j]
			m.TagCoordinates[t] = MetadataTagCoordinate{
				Cluster:      i,
				ClusterIndex: j,
			}
		}
	}
}

func (m *MetadataVectors) resolveClusterSizes() {
	for i := 0; i < len(m.Clusters); i++ {
		cluster := m.Clusters[i]
		m.ClusterSizes[i] = len(cluster)
	}
}

/*
ReadMetadataFromFile reads MetadataVectors from given json file

*/
func ReadMetadataVectorsFromFile(fileName string) (*MetadataVectors, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer finalize(file)

	decoder := json.NewDecoder(file)
	v := MetadataVectors{}
	err = decoder.Decode(&v)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
