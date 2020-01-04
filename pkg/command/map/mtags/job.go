package mtags

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
)

const supportedExtension = ".json"

type tagCollectionJob struct {
	env  *tagCollectorEnvironment
	path string
	info os.FileInfo
}

func (t tagCollectionJob) Handle() error {
	if !strings.HasSuffix(t.info.Name(), supportedExtension) {
		return nil
	}
	return t.handleJSONPath()
}

func (t *tagCollectionJob) handleJSONPath() error {
	track, err := model.ReadRawTrackFromFile(t.path)
	if err != nil {
		return err
	}

	for _, array := range track.Tags {
		err = t.handleTagArray(array)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *tagCollectionJob) handleTagArray(array []string) error {
	if len(array) != 2 {
		message := fmt.Sprintf("tag array has illegal size: %d expected: 2\n", len(array))
		return errors.New(message)
	}

	i, err := strconv.ParseInt(array[1], 10, 16)
	if err != nil {
		message := fmt.Sprintf("could not parse tag weight because of an error: %v\n", err)
		return errors.New(message)
	}

	tag := tag{
		Name:   array[0],
		Weight: i,
	}
	err = t.env.collector.register(&tag)
	if err != nil {
		return err
	}

	return nil
}

type tagCollectionJobFactory struct {
	env *tagCollectorEnvironment
}

func (t tagCollectionJobFactory) Create(path string, info os.FileInfo) (utils.Job, error) {
	return tagCollectionJob{
		env:  t.env,
		path: path,
		info: info,
	}, nil
}
