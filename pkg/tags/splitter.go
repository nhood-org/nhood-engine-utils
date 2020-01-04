package tags

import "strings"

// Splitter implements Resolver interface
// and resolves tags by splitting them and filering out common words
type Splitter struct {
}

// Resolve will split given in a slice of tags
func (s Splitter) Resolve(tag string) ([]string, error) {
	tags := strings.Split(tag, " ")
	filtered := make([]string, 0)
	for _, t := range tags {
		name := strings.ToLower(t)
		isValid := true
		isValid = isValid && nameIsNotAuxiliaryWord(name)
		isValid = isValid && nameIsNotASingleCharacter(name)
		if isValid {
			filtered = append(filtered, name)
		}
	}
	return filtered, nil
}

// NewSplitter return a pointer to a new instance of Splitter
func NewSplitter() *Splitter {
	return &Splitter{}
}
