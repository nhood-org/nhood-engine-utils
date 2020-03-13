#!/bin/sh

set -e

make build

CMD="./bin/engine-utils"
INPUT_DATA="${HOME}/Downloads/lastfm_train"

$CMD w2v $INPUT_DATA --size=15