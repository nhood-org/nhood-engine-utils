package service

import (
	"bufio"
	"os"
)

var auxiliaryWords = make(map[string]bool)

func init() {
	file, err := os.Open("./pkg/service/auxiliary_words")
	if err != nil {
		panic(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w := scanner.Text()
		auxiliaryWords[w] = true
	}
}

func nameIsNotAuxiliaryWord(name string) bool {
	return !auxiliaryWords[name]
}
