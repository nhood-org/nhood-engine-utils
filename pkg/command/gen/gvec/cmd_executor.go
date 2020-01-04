package gvec

import (
	"fmt"
	"log"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
	"github.com/spf13/cobra"
)

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	mapper := NewMapperImpl()

	filterConfig := newTagFilterImplConfig(
		args.MinimalTagWeight,
	)
	filter := newTagFilterImpl(filterConfig)

	crawlerConfig := newCrawlerConfig(
		args.Threshold,
		args.TagIrrelevanceThreshold,
		args.ClusteringrelevancePercentage,
		args.ClusteringIrrelevancePercentage,
	)

	crawler := newCrawler(crawlerConfig)

	vectorResolver := newVectorResolverImpl()

	env := generateVectorsEnvironment{
		args:               args,
		tagRelevanycMapper: mapper,
		tagFilter:          filter,
		clusterResolver:    crawler,
		vectorResolver:     vectorResolver,
	}

	tags, err := model.ReadTagsFromFile(args.TagsInput)
	if err != nil {
		panic(err)
	}
	log.Println("There are", len(*tags), "tags to process")

	tracks, err := model.ReadTracksFromFile(args.TracksInput)
	if err != nil {
		panic(err)
	}
	log.Println("There are", len(*tracks), "tracks to process")

	matrix := env.tagRelevanycMapper.resolve(tags, tracks)
	matrix = env.tagFilter.filter(matrix)
	clusters := env.clusterResolver.resolve(matrix)
	vectors := env.vectorResolver.resolve(matrix, clusters)

	err = utils.SaveToFile(vectors, args.Output)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resolved %d tags in %d clusters", matrix.Len(), len(clusters))
}
