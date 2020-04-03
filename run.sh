#!/bin/sh

set -e

make build

CMD="./bin/engine-utils"
INPUT_DATA="./test_data/"

$CMD generate-corpus $INPUT_DATA
$CMD word2vec corpus.out --size=15