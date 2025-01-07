package pocketlog_test

import (
	"hollowmatt/logger/pocketlog"
	"io"
	"os"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, io.Writer(os.Stdout))
	debugLogger.Debugf("Hello %s", "logger")
}
