package mtracks

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
	"github.com/nhood-org/nhood-engine-utils/pkg/utils"
)

const supportedExtension = ".json"

type trackMapperJob struct {
	env  *trackMapperEnvironment
	path string
	info os.FileInfo
}

func (t trackMapperJob) Handle() error {
	if !strings.HasSuffix(t.info.Name(), supportedExtension) {
		return nil
	}
	return t.mapJSONPath()
}

func (t *trackMapperJob) mapJSONPath() error {
	raw, err := model.ReadRawTrackFromFile(t.path)
	if err != nil {
		return err
	}

	similarIDs, err := t.resolveTrackSimilarIDs(raw)
	if err != nil {
		return err
	}

	tags, err := t.resolveTrackTags(raw)
	if err != nil {
		return err
	}

	track := &model.Track{
		ID:         raw.ID,
		SimilarIDs: similarIDs,
		Tags:       tags,
	}

	t.env.collector.register(track)

	return nil
}

func (t *trackMapperJob) resolveTrackSimilarIDs(raw *model.RawTrack) ([]string, error) {
	similarIDs := make([]string, 0)

	for _, id := range raw.SimilarIDs {
		similarID := fmt.Sprintf("%v", id[0])

		similarityWeight, err := strconv.ParseFloat(fmt.Sprintf("%v", id[1]), 64)
		if err != nil {
			return nil, err
		}
		if similarityWeight < t.env.args.SimilarityThreshold {
			continue
		}

		similarID = strings.TrimSpace(similarID)
		similarIDs = append(similarIDs, similarID)
	}

	return similarIDs, nil
}

func (t *trackMapperJob) resolveTrackTags(raw *model.RawTrack) ([]string, error) {
	knownTags := t.env.knownTags
	tags := make([]string, 0)

	for _, tag := range raw.Tags {
		tagName := tag[0]
		if !knownTags.Contains(tagName) {
			continue
		}

		tagWeight, err := strconv.Atoi(tag[1])
		if err != nil {
			return nil, err
		}
		if tagWeight < t.env.args.TagThreshold {
			continue
		}

		tags = append(tags, tagName)
	}

	return tags, nil
}

type trackMapperJobFactory struct {
	env *trackMapperEnvironment
}

func (t trackMapperJobFactory) Create(path string, info os.FileInfo) (utils.Job, error) {
	return trackMapperJob{
		env:  t.env,
		path: path,
		info: info,
	}, nil
}
