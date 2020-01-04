package gvec

import (
	"log"
	"sort"

	"github.com/nhood-org/nhood-engine-utils/pkg/command/gen/gvec/relevance"
)

type crawler struct {
	config *crawlerConfig
}

type crawlerConfig struct {
	Threshold                    int
	tagIrrelevanceThreshold      int
	clusterRelevancePercentage   int
	clusterIrrelevancePercentage int
}

type crawlerResolutionContext struct {
	config    *crawlerConfig
	matrix    *relevance.Matrix
	clustered map[string]bool
}

func newCrawlerConfig(
	Threshold int,
	tagIrrelevanceThreshold int,
	clusterRelevancePercentage int,
	clusterIrrelevancePercentage int,
) *crawlerConfig {
	return &crawlerConfig{
		Threshold:                    Threshold,
		tagIrrelevanceThreshold:      tagIrrelevanceThreshold,
		clusterRelevancePercentage:   clusterRelevancePercentage,
		clusterIrrelevancePercentage: clusterIrrelevancePercentage,
	}
}

func newCrawler(config *crawlerConfig) *crawler {
	return &crawler{
		config: config,
	}
}

func (c *crawler) resolve(matrix *relevance.Matrix) [][]string {
	clusters := make([][]string, 0)

	ctx := crawlerResolutionContext{
		config:    c.config,
		matrix:    matrix,
		clustered: make(map[string]bool, 0),
	}

	ctx.sortTagsByPopularityAsc()

	for i, t := range matrix.Tags() {
		if ctx.isTagClusteredAlready(t) {
			continue
		}

		cluster := map[string]bool{
			t: true,
		}

		irrelevant := ctx.resolveIrrelevantTags(t)
		cluster = appendAll(cluster, irrelevant)
		cluster = ctx.crawlAndMerge(cluster, irrelevant)

		next := make([]string, 0)
		for ct := range cluster {
			ctx.tagClustered(ct)
			next = append(next, ct)
		}
		clusters = append(clusters, next)

		if i%1000 == 0 {
			log.Println("Crawled:", i, "tags")
		}
	}
	return clusters
}

func (c *crawlerResolutionContext) sortTagsByPopularityAsc() {
	matrix := c.matrix

	tags := matrix.Tags()
	sort.Slice(tags, func(i, j int) bool {
		ti := tags[i]
		si := matrix.Weight(ti)
		tj := tags[j]
		sj := matrix.Weight(tj)
		return si < sj
	})

	matrix.SetTags(tags)
}

func (c *crawlerResolutionContext) resolveIrrelevantTags(tag string) map[string]bool {
	irrelevant := make(map[string]bool)
	for _, t := range c.matrix.Tags() {
		r := c.matrix.Relevance(tag, t)
		if c.isTagClusteredAlready(t) {
			continue
		}
		if r <= c.config.tagIrrelevanceThreshold && tag != t {
			irrelevant[t] = true
		}
	}
	return irrelevant
}

func (c *crawlerResolutionContext) resolveRelevantTags(tag string) map[string]bool {
	relevant := make(map[string]bool)
	for _, t := range c.matrix.Tags() {
		r := c.matrix.Relevance(tag, t)
		if c.isTagClusteredAlready(t) {
			continue
		}
		if r >= c.config.Threshold && tag != t {
			relevant[t] = true
		}
	}
	return relevant
}

func (c *crawlerResolutionContext) tagClustered(tag string) {
	c.clustered[tag] = true
}

func (c *crawlerResolutionContext) isTagClusteredAlready(tag string) bool {
	_, contains := c.clustered[tag]
	return contains
}

func (c *crawlerResolutionContext) crawlAndMerge(cluster map[string]bool, irrelevant map[string]bool) map[string]bool {
	for t := range irrelevant {
		tRelevant := c.resolveRelevantTags(t)
		commonRelevant := common(cluster, tRelevant)
		if isGraterEqualThan(commonRelevant, cluster, c.config.clusterRelevancePercentage) {
			continue
		}

		tIrrelevant := c.resolveIrrelevantTags(t)
		commonIrrelevant := common(cluster, tIrrelevant)
		if isGraterEqualThan(commonIrrelevant, cluster, c.config.clusterIrrelevancePercentage) {
			cluster = commonIrrelevant
			cluster[t] = true
		}
	}
	return cluster
}

func isGraterEqualThan(m1 map[string]bool, m2 map[string]bool, percentage int) bool {
	m1l := len(m1)
	m2l := len(m2)
	return (m1l*100)/m2l >= percentage
}

func common(m1 map[string]bool, m2 map[string]bool) map[string]bool {
	common := make(map[string]bool)
	for v := range m2 {
		_, contains := m1[v]
		if contains {
			common[v] = true
		}
	}
	return common
}

func appendAll(m1 map[string]bool, m2 map[string]bool) map[string]bool {
	for v := range m2 {
		m1[v] = true
	}
	return m1
}
