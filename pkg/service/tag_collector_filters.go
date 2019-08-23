package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var auxiliaryWords = make(map[string]bool)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	auxiliaryWordsFile := "./auxiliary_words"
	if !strings.Contains(pwd, "/pkg/service") {
		auxiliaryWordsFile = "./pkg/service/auxiliary_words"
	}

	file, err := os.Open(auxiliaryWordsFile)
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

func nameIsNotASingleCharacter(name string) bool {
	return len(name) != 1
}
