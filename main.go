package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

/*
TestDataDirectory defines an entrypoint of test data

*/
var TestDataDirectory = "./test-data"

/*
TestDataJSONExtension defines supported file type extension

*/
var TestDataJSONExtension = ".json"

func main() {
	filepath.Walk(TestDataDirectory, handlePath)
}

func handlePath(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(info.Name(), TestDataJSONExtension) {
		go handleJSONPath(path)
	}
	return nil
}

func handleJSONPath(path string) {
	t, _ := model.ReadTrack(path)
	println(t.ToString())
}
