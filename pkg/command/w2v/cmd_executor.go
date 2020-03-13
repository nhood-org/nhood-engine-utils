package w2v

import (
	"os"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
	"github.com/spf13/cobra"
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

	generateCorpus(args)
	generateVectors(args)
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
