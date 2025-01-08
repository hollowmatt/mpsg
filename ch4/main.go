package main

import (
	"hollowmatt/logger/pocketlog"
	"os"
	"time"
)

// Example of how to implement the logger
func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout))

	lgr.Infof("This is for your information only.")
	lgr.Errorf("An error has occured in %s", "main")
	lgr.Debugf("line 14 of main")
	lgr.Infof("The time is now %s", time.Now())
}
