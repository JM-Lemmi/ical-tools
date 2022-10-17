package main

import (
	"os"

	ics "github.com/arran4/golang-ical"
	log "github.com/sirupsen/logrus"
)

var version = "0.1.0"

func main() {
	log.Info("Welcome to ical-tools, version " + version)
	log.SetLevel(log.DebugLevel)
	log.Debug("Debug log is enabled") // only shows if Debug is actually enabled

	//start module from os.Args[1]

	var calendar1 *ics.Calendar
	file1, _ := os.Open("./ical-1.ics")
	calendar1, _ = ics.ParseCalendar(file1)

	var calendar2 *ics.Calendar
	file2, _ := os.Open("./ical-2.ics")
	calendar2, _ = ics.ParseCalendar(file2)

	compare(calendar1, calendar2)
}
