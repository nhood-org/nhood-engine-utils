package gvec

import (
	"testing"

	"github.com/nhood-org/nhood-engine-utils/pkg/command/gen/gvec/relevance"
	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestClustersAreResolved(t *testing.T) {
	a := assert.New(t)

	tags := &model.Tags{
		"jazz":   true,
		"blues":  true,
		"rock":   true,
		"metal":  true,
		"UK":     true,
		"USA":    true,
		"Polish": true,
		"80":     true,
		"90":     true,
		"00":     true,
	}

	matrix := relevance.NewMatrix(tags)
	registerTags(matrix, "jazz", "USA", "80")
	registerTags(matrix, "jazz", "USA", "90")
	registerTags(matrix, "jazz", "USA", "00")
	registerTags(matrix, "jazz", "Polish", "80")
	registerTags(matrix, "jazz", "Polish", "90")
	registerTags(matrix, "jazz", "Polish", "00")

	registerTags(matrix, "blues", "USA", "90")
	registerTags(matrix, "blues", "USA", "00")
	registerTags(matrix, "blues", "UK", "80")
	registerTags(matrix, "blues", "UK", "90")

	registerTags(matrix, "rock", "UK", "80")
	registerTags(matrix, "rock", "UK", "90")
	registerTags(matrix, "rock", "UK", "00")
	registerTags(matrix, "rock", "USA", "80")
	registerTags(matrix, "rock", "USA", "90")
	registerTags(matrix, "rock", "USA", "00")
	registerTags(matrix, "rock", "Polish", "80")
	registerTags(matrix, "rock", "Polish", "90")
	registerTags(matrix, "rock", "Polish", "00")

	registerTags(matrix, "metal", "USA", "80")
	registerTags(matrix, "metal", "USA", "90")
	registerTags(matrix, "metal", "USA", "00")
	registerTags(matrix, "metal", "Polish", "90")
	registerTags(matrix, "metal", "Polish", "00")

	crawlerConfig := newCrawlerConfig(
		thresholdDefault,
		tagIrrelevanceThresholdDefault,
		clusteringRelevancePercentageDefault,
		clusteringIrrelevancePercentageDefault,
	)
	crawler := newCrawler(crawlerConfig)
	clusters := crawler.resolve(matrix)

	a.Len(clusters, 3)

	c0 := clusters[0]
	a.Contains(c0, "blues")
	a.Contains(c0, "metal")
	a.Contains(c0, "jazz")
	a.Contains(c0, "rock")

	c1 := clusters[1]
	a.Contains(c1, "UK")
	a.Contains(c1, "Polish")
	a.Contains(c1, "USA")

	c2 := clusters[2]
	a.Contains(c2, "80")
	a.Contains(c2, "90")
	a.Contains(c2, "00")
}

func registerTags(matrix *relevance.Matrix, tags ...string) {
	for _, t1 := range tags {
		for _, t2 := range tags {
			matrix.Increment(t1, t2)
		}
	}
}
