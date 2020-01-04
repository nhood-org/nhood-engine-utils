package utils

import (
	"os"
)

/*
OutputFile is a simple wrapper struct for os.File

*/
type OutputFile struct {
	f *os.File
}

/*
NewOutputFile returns a new instance of OutputFile

*/
func NewOutputFile(fileName string) (*OutputFile, error) {
	info, err := os.Stat(fileName)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if info != nil {
		err = os.Remove(fileName)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &OutputFile{
		f: file,
	}, nil
}

/*
Append adds a new line to the output file

*/
func (o *OutputFile) Append(line string) error {
	_, err := o.f.Write([]byte(line + "\n"))
	if err != nil {
		return err
	}
	return nil
}

/*
Close closes an output file

*/
func (o *OutputFile) Close() error {
	return o.f.Close()
}
