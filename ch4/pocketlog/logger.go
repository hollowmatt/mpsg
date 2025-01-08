package pocketlog

import (
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns a shiny new logger, ready to log at the threshold you desire
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}
