package arguments

import (
	"errors"
	"os"
)

/*
Arguments contains all possible application arguments

*/
type Arguments struct {
	Root string
}

/*
ResolveArguments will parse all arguments and return as Arguments structure

*/
func ResolveArguments(args []string) (*Arguments, error) {
	if len(args) == 0 {
		return nil, errors.New("Directory argument is required")
	}

	root := args[0]
	if _, err := os.Stat(root); err == nil {

	} else if os.IsNotExist(err) {
		return nil, errors.New("Directory '" + root + "' does not exist")
	} else {
		return nil, errors.New("Could not check '" + root + "' directory")
	}

	return &Arguments{
		Root: root,
	}, nil
}
