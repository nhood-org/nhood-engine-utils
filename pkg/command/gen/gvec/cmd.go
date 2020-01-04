package gvec

import (
	"github.com/spf13/cobra"
)

const tagsInputFlagName = "tags"
const tracksInputFlagName = "tracks"
const outputFlagName = "output"
const minimalTagWeightFlagName = "min-tag-weight"
const thresholdFlagName = "tag-relevance-th"
const tagIrrelevanceThresholdFlagName = "tag-irrelevance-th"
const clusteringRelevancePercentageFlagName = "cluster-relevance"
const clusteringIrrelevancePercentageFlagName = "cluster-irrelevance"

const tagsInputDefault = "tags.out.json"
const tracksInputDefault = "tracks.out.json"
const outputDefault = "vectors.out.json"

// Last best results execution:
const minimalTagWeightDefault = 1000
const thresholdDefault = 100
const tagIrrelevanceThresholdDefault = 0
const clusteringRelevancePercentageDefault = 50
const clusteringIrrelevancePercentageDefault = 50

/*
NewCommand returns an instance of a cobra.Command
implementing a track metadata vector generation operations

*/
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gvec",
		Short: "Generate track metadata vectors",
		Long:  `gvec is for resolution of track metadata vectors from within the input track data structure.`,
		Args:  cobra.MinimumNArgs(0),
		Run:   execute,
	}
	cmd.Flags().StringP(tagsInputFlagName, "i", tagsInputDefault, "tags input file")
	cmd.Flags().StringP(tracksInputFlagName, "t", tracksInputDefault, "tracks input file")
	cmd.Flags().StringP(outputFlagName, "o", outputDefault, "output file")

	cmd.Flags().Uint(minimalTagWeightFlagName, minimalTagWeightDefault, `minimal tag weight.
	Tags with occurrence weight lower that given minimal tag weight will not be clustered`)

	cmd.Flags().Uint(thresholdFlagName, thresholdDefault, `tag relevance threshold. 
	Tags are considered relevant if those were used together within a single track more times than the given threshold`)

	cmd.Flags().Uint(tagIrrelevanceThresholdFlagName, tagIrrelevanceThresholdDefault, `tag irrelevance threshold. 
	Tags are considered relevant if those were used together within a single track more times than the given threshold`)

	cmd.Flags().Uint(clusteringRelevancePercentageFlagName, clusteringRelevancePercentageDefault, `relevance percentage.
	During the course of clustering tags are compared by the amount of common relevant tags. 
	It a number of common relevant tags is below the given relevance percentage of all of its relevant tags
	the tags are considered relevant for clustering`)

	cmd.Flags().Uint(clusteringIrrelevancePercentageFlagName, clusteringIrrelevancePercentageDefault, `irrelevance percentage.
	During the course of clustering tags compared by the amount of common irrelevant tags. 
	It a number of common irrelevant tags exceeds the given irrelevance percentage of all of its irrelevant tags
	the tags are considered relevant for clustering`)

	return cmd
}
