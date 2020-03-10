package w2v

import (
	"fmt"
	"os"
	"sync"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
	"github.com/spf13/cobra"
)

type job struct {
	out   chan model.RawTrack
	outWg *sync.WaitGroup
	path  string
}

func (j job) Handle() error {
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
	}, nil
}

func execute(cmd *cobra.Command, cmdArgs []string) {
	args, err := resolveArguments(cmd, cmdArgs)
	if err != nil {
		panic(err)
	}

	defaultPoolSize := 5

	var outWg sync.WaitGroup
	out := make(chan model.RawTrack)

	jobFactory := jobFactory{
		outWg: &outWg,
		out:   out,
	}

	f, err := os.OpenFile(args.Output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	go func() {
		for {
			o := <-out

			s := "track of ID " + o.ID + " is tagged as "
			for _, t := range o.Tags {
				s += t[0] + " "
			}
			s += "and similar to "
			for _, id := range o.SimilarIDs {
				s += id[0].(string) + " "
			}
			fmt.Fprint(f, s)

			outWg.Done()
		}
	}()

	pathWalker := utils.NewPathWalker(defaultPoolSize, args.Root, jobFactory)
	pathWalker.Execute()

	outWg.Wait()
}
