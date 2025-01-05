package pocketlog

import "fmt"

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

	_, _ = fmt.Printf("DEBUG: "+format+"\n", args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	_, _ = fmt.Printf("INFO: "+format+"\n", args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}

	_, _ = fmt.Printf("WARN: "+format+"\n", args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	_, _ = fmt.Printf("ERROR: "+format+"\n", args...)
}

func (l *Logger) Fatalf(format string, args ...any) {
	if l.threshold > LevelFatal {
		return
	}

	_, _ = fmt.Printf("FATAL: "+format+"\n", args...)
}
