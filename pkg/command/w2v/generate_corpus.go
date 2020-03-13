package w2v

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
)

func generateCorpus(args word2VecCommandArguments) {
	var outWg sync.WaitGroup
	out := make(chan model.RawTrack)

	jobFactory := jobFactory{
		outWg: &outWg,
		out:   out,
	}

	tracksFileName := getTracksFileName(args.Output)
	fTracks, err := os.Create(tracksFileName)
	if err != nil {
		panic(err)
	}
	defer fTracks.Close()

	corpusFileName := getCorpusFileName(args.Output)
	fCorpus, err := os.Create(corpusFileName)
	if err != nil {
		panic(err)
	}
	defer fCorpus.Close()

	go func() {
		for {
			o := <-out
			fmt.Fprintf(fTracks, "%s, %s, %s\n", o.ID, o.Artist, o.Title)

			var s string
			for _, t := range o.Tags {
				s = appendTag(args, o.ID, s, t)
			}
			for _, id := range o.SimilarIDs {
				s = appendSimilarID(args, o.ID, s, id)
			}

			if s != "" {
				fmt.Fprint(fCorpus, s)
			}

			outWg.Done()
		}
	}()

	pathWalker := utils.NewPathWalker(defaultPoolSize, args.Root, jobFactory)
	pathWalker.Execute()

	outWg.Wait()
}

func appendTag(args word2VecCommandArguments, trackID string, s string, tag []string) string {
	tagWeight, err := strconv.Atoi(tag[1])
	if err != nil {
		return s
	}
	if tagWeight < args.ThresholdTag {
		return s
	}
	return s + fmt.Sprintf("%s tagged as %s\n", trackID, tag[0])
}

func appendSimilarID(args word2VecCommandArguments, trackID string, s string, id []interface{}) string {
	similarityWeight, err := strconv.ParseFloat(fmt.Sprintf("%v", id[1]), 64)
	if err != nil {
		return s
	}
	if similarityWeight < args.ThresholdSimilarity {
		return s
	}
	return s + fmt.Sprintf("%s similar to %s\n", trackID, id[0].(string))
}
