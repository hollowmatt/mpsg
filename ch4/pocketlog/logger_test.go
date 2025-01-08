package pocketlog_test

import (
	"hollowmatt/logger/pocketlog"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello %s", "logger")
}

type testWriter struct {
	contents string
}

func (t *testWriter) Write(p []byte) (n int, err error) {
	t.contents += string(p)
	return len(p), nil
}

const (
	debugMsg = "This is a debug level message"
	infoMsg  = "This is an info level message"
	errMsg   = "This is an error level message"
	fatalMsg = "This is a fatal level message"
)

func TestLogger_Levels(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: debugMsg + "\n" + infoMsg + "\n" + errMsg + "\n" + fatalMsg + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: infoMsg + "\n" + errMsg + "\n" + fatalMsg + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: errMsg + "\n" + fatalMsg + "\n",
		},
		"fatal": {
			level:    pocketlog.LevelFatal,
			expected: fatalMsg + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			testWriter := &testWriter{}
			lgr := pocketlog.New(tc.level, pocketlog.WithOutput(testWriter))
			lgr.Debugf(debugMsg)
			lgr.Infof(infoMsg)
			lgr.Errorf(errMsg)
			lgr.Fatalf(fatalMsg)

			if testWriter.contents != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, testWriter.contents)
			}
		})
	}
}
