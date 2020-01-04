package tags

// Resolver defines an interface resolving a slice of tags from a single string
type Resolver interface {
	Resolve(tag string) ([]string, error)
}
