package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

func main() {
	var root = "./test-data"
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".json") {
			go func(path string) {
				t, _ := model.ReadTrack(path)
				println(t.ToString())
			}(path)
		}
		return nil
	})
}
