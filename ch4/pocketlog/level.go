package pocketlog

import (
	"fmt"
	"io"
	"os"
)

var (
	Stdout io.Writer = os.Stdout
	Stderr io.Writer = os.Stderr
)

// used to represent the log levels
type Level byte

const (
	//lowest level of logging, used for development and debugging
	LevelDebug Level = iota
	//level of logging used for informational purposes (no errors occurred)
	LevelInfo
	//level indicating a warning event occurred
	LevelWarn
	//level indicating an error has occurred in the system, may be recoverable
	LevelError
	//level indicating a fatal event occurred, non-recoverable
	LevelFatal
)

func (l *Logger) Debugf(format string, args ...any) {
	// making sure we can write to the output
	if l.output == nil {
		l.output = Stdout
	}

	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Fprintf(l.output, "DUBUG: "+format, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = Stdout
	}

	if l.threshold > LevelInfo {
		return
	}

	_, _ = fmt.Fprintf(l.output, "INFO: "+format, args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	if l.output == nil {
		l.output = Stdout
	}

	if l.threshold > LevelWarn {
		return
	}

	_, _ = fmt.Fprintf(l.output, "WARN: "+format+"\n", args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = Stderr
	}

	if l.threshold > LevelError {
		return
	}

	_, _ = fmt.Fprintf(l.output, "ERROR: "+format+"\n", args...)
}

func (l *Logger) Fatalf(format string, args ...any) {
	if l.output == nil {
		l.output = Stderr
	}

	if l.threshold > LevelFatal {
		return
	}

	_, _ = fmt.Fprintf(l.output, "FATAL: "+format+"\n", args...)
}
