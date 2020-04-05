#!/bin/sh

set -e

make build

CMD="./bin/engine-utils"

INPUT_DATA=$1
if [ -z "$INPUT_DATA" ]
  then
    INPUT_DATA="./test_data/"
fi

TRACKS_OUT="tracks.out"
TAGS_CORPUS_OUT="tags-corpus.out"
TAGS_VECTORS_OUT="tags-vectors.out"

$CMD map-tracks $INPUT_DATA \
  --output=$TRACKS_OUT
$CMD generate-corpus $TRACKS_OUT \
  --mode=TAGS \
  --output=$TAGS_CORPUS_OUT
$CMD glove $TAGS_CORPUS_OUT \
  --output=$TAGS_VECTORS_OUT
$CMD generate-vectors $TRACKS_OUT $TAGS_VECTORS_OUT \
  --output=vectors.out.csv
