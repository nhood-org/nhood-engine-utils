package tags

import (
	"regexp"
	"strings"
)

const nonAlphanumericRegex = "[^a-zA-Z0-9]+"

// Sanitizer implements Resolver interface
// and resolves tags by unifying them
// by trimming and non-alphanumeric character removal
type Sanitizer struct {
}

// Resolve will split given in a slice of tags
func (s Sanitizer) Resolve(tag string) ([]string, error) {
	reg, err := regexp.Compile(nonAlphanumericRegex)
	if err != nil {
		return nil, err
	}
	result := reg.ReplaceAllString(tag, "")
	result = strings.ToLower(result)
	return []string{result}, nil
}

// NewSanitizer return a pointer to a new instance of Sanitizer
func NewSanitizer() *Sanitizer {
	return &Sanitizer{}
}
