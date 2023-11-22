package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type stdReader struct {
	reader io.Reader
}

func newPipeReader() *stdReader {
	return &stdReader{
		reader: os.Stdin,
	}
}

func (r *stdReader) Read() (string, error) {
	var input string
	scanner := bufio.NewScanner(bufio.NewReader(r.reader))
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return input, err
		}
		input += fmt.Sprintln(scanner.Text())
	}
	return input, nil
}
