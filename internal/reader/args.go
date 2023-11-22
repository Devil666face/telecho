package reader

import (
	"errors"
)

var errArgsNotSet = errors.New("you not set any args")

type argsReader struct {
	args []string
}

func newArgsReader(_args []string) *argsReader {
	return &argsReader{
		args: _args,
	}
}

func (r *argsReader) Read() (string, error) {
	var input string
	if len(r.args) == 0 {
		return input, errArgsNotSet
	}
	for _, a := range r.args {
		input += a + " "
	}
	return input, nil
}
