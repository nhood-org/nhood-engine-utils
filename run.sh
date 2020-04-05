#!/bin/sh

set -e

make build

CMD="./bin/engine-utils"

INPUT_DATA=$1
if [ -z "$INPUT_DATA" ]
  then
    INPUT_DATA="./test_data/"
fi

$CMD map-tracks $INPUT_DATA
$CMD generate-corpus ./tracks.out --mode=TAGS
$CMD glove corpus.out
