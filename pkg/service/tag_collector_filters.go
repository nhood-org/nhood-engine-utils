package service

import "github.com/nhood-org/nhood-engine-utils/pkg/utils"

const defaultResultTagCountThreshold = 5000

var auxiliaryWords = make(map[string]bool)

func init() {
	initAuxiliaryWords("a", "an", "the", "is",
		"was", "be", "been", "have", "has", "had", "not",
		"do", "doing", "done", "for", "from", "and", "nor", "but", "or", "yet", "so", "with", "without",
		"who", "whom", "whomever", "whose", "which", "why", "when", "what",
		"this", "these", "that", "those",
		"i", "you", "he", "she", "it", "me", "my", "your", "his", "her", "mine", "yours", "am", "are",
		"at", "by", "in", "into", "on", "of", "to", "all", "no", "none", "out", "-")
}

func initAuxiliaryWords(words ...string) {
	for _, w := range words {
		auxiliaryWords[w] = true
	}
}

func nameIsNotAuxiliaryWord(name string) bool {
	return !auxiliaryWords[name]
}

func hasSufficientCount(ma *utils.MovingAverage) bool {
	return ma.Count() >= defaultResultTagCountThreshold
}
