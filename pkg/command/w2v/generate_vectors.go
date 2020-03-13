package w2v

import (
	"os"

	"github.com/ynqa/wego/pkg/builder"
	"github.com/ynqa/wego/pkg/model/word2vec"
)

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
