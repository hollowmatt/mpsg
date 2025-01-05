package pocketlog

type Logger struct {
	threshold Level
}

// New returns a shiny new logger, ready to log at the threshold you desire
func New(threshold Level) *Logger {
	return &Logger{threshold: threshold}
}
