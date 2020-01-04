package gvec

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type generateVectorsCommandArguments struct {
	Output                          string
	TagsInput                       string
	TracksInput                     string
	MinimalTagWeight                int
	Threshold                       int
	TagIrrelevanceThreshold         int
	ClusteringrelevancePercentage   int
	ClusteringIrrelevancePercentage int
}

func resolveArguments(cmd *cobra.Command, args []string) (*generateVectorsCommandArguments, error) {
	minimalTagWeight, err := strconv.Atoi(cmd.Flag(minimalTagWeightFlagName).Value.String())
	if err != nil {
		return nil, errors.New(minimalTagWeightFlagName + " flag is invalid")
	}
	if minimalTagWeight < 0 {
		return nil, errors.New(minimalTagWeightFlagName + " flag must be grater that 0")
	}

	Threshold, err := strconv.Atoi(cmd.Flag(thresholdFlagName).Value.String())
	if err != nil {
		return nil, errors.New(thresholdFlagName + " flag is invalid")
	}
	if Threshold < 0 {
		return nil, errors.New(thresholdFlagName + " flag must be grater that 0")
	}

	tagIrrelevanceThreshold, err := strconv.Atoi(cmd.Flag(tagIrrelevanceThresholdFlagName).Value.String())
	if err != nil {
		return nil, errors.New(tagIrrelevanceThresholdFlagName + " flag is invalid")
	}
	if tagIrrelevanceThreshold < 0 {
		return nil, errors.New(tagIrrelevanceThresholdFlagName + " flag must be grater that 0")
	}

	clusteringrelevancePercentage, err := strconv.Atoi(cmd.Flag(clusteringRelevancePercentageFlagName).Value.String())
	if err != nil {
		return nil, errors.New(clusteringRelevancePercentageFlagName + " flag is invalid")
	}
	if clusteringrelevancePercentage < 0 {
		return nil, errors.New(clusteringRelevancePercentageFlagName + " flag must be grater that 0")
	}
	if clusteringrelevancePercentage > 100 {
		return nil, errors.New(clusteringRelevancePercentageFlagName + " flag must not be grater that 100")
	}

	clusteringIrrelevancePercentage, err := strconv.Atoi(cmd.Flag(clusteringIrrelevancePercentageFlagName).Value.String())
	if err != nil {
		return nil, errors.New(clusteringIrrelevancePercentageFlagName + " flag is invalid")
	}
	if clusteringIrrelevancePercentage < 0 {
		return nil, errors.New(clusteringIrrelevancePercentageFlagName + " flag must be grater that 0")
	}
	if clusteringIrrelevancePercentage > 100 {
		return nil, errors.New(clusteringIrrelevancePercentageFlagName + " flag must not be grater that 100")
	}

	tagsInput := cmd.Flag(tagsInputFlagName).Value.String()
	if _, err := os.Stat(tagsInput); err == nil {

	} else if os.IsNotExist(err) {
		return nil, errors.New("file '" + tagsInput + "' does not exist")
	} else {
		return nil, errors.New("could not check '" + tagsInput + "' input file")
	}

	tracksInput := cmd.Flag(tracksInputFlagName).Value.String()
	if _, err := os.Stat(tracksInput); err == nil {

	} else if os.IsNotExist(err) {
		return nil, errors.New("file '" + tracksInput + "' does not exist")
	} else {
		return nil, errors.New("could not check '" + tracksInput + "' input file")
	}

	output := cmd.Flag(outputFlagName).Value.String()

	return &generateVectorsCommandArguments{
		Output:                          output,
		TagsInput:                       tagsInput,
		TracksInput:                     tracksInput,
		MinimalTagWeight:                minimalTagWeight,
		Threshold:                       Threshold,
		TagIrrelevanceThreshold:         tagIrrelevanceThreshold,
		ClusteringrelevancePercentage:   clusteringrelevancePercentage,
		ClusteringIrrelevancePercentage: clusteringIrrelevancePercentage,
	}, nil
}
