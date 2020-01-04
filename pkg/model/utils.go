package model

import "os"

func finalize(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}
