package utils

import (
	"encoding/json"
	"os"
)

/*
SaveToFile saves metadata vectors to json file

*/
func SaveToFile(e interface{}, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer finalize(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(e)
	if err != nil {
		return err
	}

	return nil
}

func finalize(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}
