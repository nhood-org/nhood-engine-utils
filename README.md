[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Version](https://img.shields.io/badge/version-v0.0.4-blue.svg?maxAge=2592000)](https://github.com/nhood-org/nhood-engine-utils/releases/tag/v0.0.4)
[![CircleCI](https://circleci.com/gh/nhood-org/nhood-engine-utils.svg?style=shield)](https://circleci.com/gh/nhood-org/nhood-engine-utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/nhood-org/nhood-engine-utils)](https://goreportcard.com/report/github.com/nhood-org/nhood-engine-utils)

# nhood-engine-utils
A simple Go application for processing of test input data

## Pre-requisites

- Go
- Make

## Build

In order to build the project use the following make command:

```bash
make clean build
```

## Test

In order to test the project use the following make command:

```bash
make clean test
```

## Run

In order to run the project locally use the following make command:

```bash
make clean run
```

## Usage

For usage hits see `run.sh` script as reference or use:

```bash
./bin/engine-utils --help
```

## CI/CD

Project is continuously integrated with `circleCi` pipeline that link to which may be found [here](https://circleci.com/gh/nhood-org/workflows/nhood-engine-utils)

Pipeline is fairly simple:

1. Build and test project

Configuration of CI is implemented in `.circleci/config.yml`.

## Versioning

In order to release version, execute the following script:

```bash
export CIRCLE_CI_USER_TOKEN=<CIRCLE_CI_USER_TOKEN>
export NEW_VERSION=<NEW_VERSION>
make trigger-circle-ci-release
```

## License

`nhood-engine-utils` is released under the MIT license:
- https://opensource.org/licenses/MIT
