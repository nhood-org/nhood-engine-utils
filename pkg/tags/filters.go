package tags

import (
	"strings"
)

var auxiliaryWords = make(map[string]bool)

func init() {
	auxiliaryWordsSlice := strings.Split(auxiliaryWordsRawList, "\n")
	for _, w := range auxiliaryWordsSlice {
		auxiliaryWords[w] = true
	}
}

func nameIsNotAuxiliaryWord(name string) bool {
	return !auxiliaryWords[name]
}

func nameIsNotASingleCharacter(name string) bool {
	return len(name) != 1
}
