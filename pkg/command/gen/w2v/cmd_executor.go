package w2v

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/ynqa/wego/pkg/builder"
	"github.com/ynqa/wego/pkg/model/word2vec"
)

const defaultPoolSize = 8

type job struct {
	out   chan model.RawTrack
	outWg *sync.WaitGroup
	path  string
	info  os.FileInfo
}

func (j job) Handle() error {
	if j.info.IsDir() {
		return nil
	}

	raw, err := model.ReadRawTrackFromFile(j.path)
	if err != nil {
		return err
	}

	j.outWg.Add(1)
	j.out <- *raw

	return nil
}

type jobFactory struct {
	out   chan model.RawTrack
	outWg *sync.WaitGroup
}

func (f jobFactory) Create(path string, info os.FileInfo) (utils.Job, error) {
	return job{
		out:   f.out,
		outWg: f.outWg,
		path:  path,
		info:  info,
	}, nil
}

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	createCorpus(args)
	generateVectors(args)
}

func createCorpus(args word2VecCommandArguments) {
	var outWg sync.WaitGroup
	out := make(chan model.RawTrack)

	jobFactory := jobFactory{
		outWg: &outWg,
		out:   out,
	}

	tracksFileName := getTracksFileName(args.Output)
	fTracks, err := os.Open(tracksFileName)
	if err != nil {
		panic(err)
	}
	defer fTracks.Close()

	corpusFileName := getCorpusFileName(args.Output)
	fCorpus, err := os.Open(corpusFileName)
	if err != nil {
		panic(err)
	}
	defer fCorpus.Close()

	go func() {
		for {
			o := <-out
			fmt.Fprintf(fTracks, "%s, %s, %s", o.ID, o.Artist, o.Title)

			var s string
			for _, t := range o.Tags {
				appendTag(args, o.ID, &s, t)
			}
			for _, id := range o.SimilarIDs {
				appendSimilarID(args, o.ID, &s, id)
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

func generateVectors(args word2VecCommandArguments) {
	m, err := builder.NewWord2vecBuilder().
		Dimension(args.Size).
		Window(5).
		Model(word2vec.CBOW).
		Optimizer(word2vec.NEGATIVE_SAMPLING).
		NegativeSampleSize(5).
		Verbose().
		Build()
	if err != nil {
		panic(err)
	}

	corpusFileName := getCorpusFileName(args.Output)
	f, err := os.Open(corpusFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = m.Train(f)
	if err != nil {
		panic(err)
	}

	out := getVectorsFileName(args.Output)
	m.Save(out)
}

func appendTag(args word2VecCommandArguments, trackID string, s *string, tag []string) {
	tagWeight, err := strconv.Atoi(tag[1])
	if err != nil {
		return
	}
	if tagWeight < args.ThresholdTag {
		return
	}
	modified := *s + fmt.Sprintf("%s tagged as %s\n", trackID, tag[0])
	s = &modified
}

func appendSimilarID(args word2VecCommandArguments, trackID string, s *string, id []interface{}) {
	similarityWeight, err := strconv.ParseFloat(fmt.Sprintf("%v", id[1]), 64)
	if err != nil {
		return
	}
	if similarityWeight < args.ThresholdSimilarity {
		return
	}
	modified := *s + fmt.Sprintf("%s similar to %s\n", trackID, id[0].(string))
	s = &modified
}

func getCorpusFileName(output string) string {
	return "corpus_" + output
}

func getTracksFileName(output string) string {
	return "tracks_" + output
}

func getVectorsFileName(output string) string {
	return "vectors_" + output
}
