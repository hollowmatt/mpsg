package pocketlog

type Logger struct {
	threshold Level
}

func New(threshold Level) *Logger {
	return &Logger{threshold: threshold}
}
