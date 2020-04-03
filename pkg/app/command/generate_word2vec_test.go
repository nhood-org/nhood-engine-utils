package command

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/nhood-org/nhood-engine-utils/pkg/adapters/model"
	"github.com/stretchr/testify/require"
)

func Test_GenerateWord2VecVectorsHandler_Handle(t *testing.T) {
	resolver := model.NewWord2VecVectorsResolverMock()
	handler := NewGenerateWord2VecVectorsCommandHandler(resolver)

	inContent := "CONTENT"
	in := bytes.NewBuffer([]byte(inContent))
	out := new(bytes.Buffer)

	cmd := GenerateWord2VecVectorsCmd{
		Corpus: in,
		Output: out,
	}

	err := handler.Handle(cmd)
	require.NoError(t, err)

	outBytes, err := ioutil.ReadAll(out)
	require.NoError(t, err)

	require.Equal(t, []byte(inContent), outBytes)
}
