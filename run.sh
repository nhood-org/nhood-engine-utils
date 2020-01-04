#!/bin/sh

set -e

make build

CMD="./bin/engine-utils"
INPUT_DATA="${HOME}/Downloads/lastfm_train"

# $CMD mtags $INPUT_DATA \
#     --threshold 500

# $CMD mtracks $INPUT_DATA \
#     --similarity-threshold 0.05 \
#     --tag-threshold 0

# $CMD gvec \
#     --tag-relevance-th=100 \
#     --tag-irrelevance-th=100 \
#     --min-tag-weight=10 \
#     --cluster-relevance=70 \
#     --cluster-irrelevance=70

$CMD gtracks ./tracks.out.json
