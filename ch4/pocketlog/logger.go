package pocketlog

import "io"

type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns a shiny new logger, ready to log at the threshold you desire
// The default output is Stdout.
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}
