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
	if l.threshold > LevelDebug {
		return
	}

	l.logf("DEBUG: "+format, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	l.logf("INFO: "+format, args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}

	l.logf("WARN: "+format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	l.logf("ERROR: "+format, args...)
}

func (l *Logger) Fatalf(format string, args ...any) {
	if l.threshold > LevelFatal {
		return
	}

	l.logf("FATAL: "+format, args...)
}

func (l *Logger) logf(format string, args ...any) {
	//If no output specified, use Stdout
	if l.output == nil {
		l.output = Stdout
	}
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
