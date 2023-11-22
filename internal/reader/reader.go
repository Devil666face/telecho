package reader

import "os"

type Readable interface {
	Read() (string, error)
}

func New(args []string) Readable {
	if isInputFromPipe() {
		return newPipeReader()
	}
	return newArgsReader(args)
}

func isInputFromPipe() bool {
	if fi, err := os.Stdin.Stat(); err == nil {
		return fi.Mode()&os.ModeCharDevice == 0
	}
	return false
}
