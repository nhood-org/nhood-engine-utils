package w2v

import (
	"fmt"
	"os"
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

func createCorpus(args *word2VecCommandArguments) {
	var outWg sync.WaitGroup
	out := make(chan model.RawTrack)

	jobFactory := jobFactory{
		outWg: &outWg,
		out:   out,
	}

	fName := getCorpusFileName(args.Output)
	f, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	go func() {
		for {
			o := <-out

			var s string
			for _, t := range o.Tags {
				s += fmt.Sprintf("%s tagged as %s\n", o.ID, t[0])
			}
			for _, id := range o.SimilarIDs {
				s += fmt.Sprintf("%s similar to %s\n", o.ID, id[0].(string))
			}

			if s != "" {
				fmt.Fprint(f, s)
			}

			outWg.Done()
		}
	}()

	pathWalker := utils.NewPathWalker(defaultPoolSize, args.Root, jobFactory)
	pathWalker.Execute()

	outWg.Wait()
}

func generateVectors(args *word2VecCommandArguments) {
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

	fName := getCorpusFileName(args.Output)
	f, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = m.Train(f)
	if err != nil {
		panic(err)
	}

	m.Save(args.Output)
}

func getCorpusFileName(output string) string {
	return "corpus_" + output
}
