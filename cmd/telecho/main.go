package main

import (
	"os"

	"github.com/Devil666face/telecho/internal/cmd"
)

var vers string

func main() {
	c := cmd.New(vers)
	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
