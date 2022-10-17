package main

import (
	"fmt"
	"os"

	ics "github.com/arran4/golang-ical"
	log "github.com/sirupsen/logrus"
)

var version = "0.1.0"

func main() {
	log.Info("Welcome to ical-tools, version " + version)
	log.SetLevel(log.InfoLevel)
	log.Debug("Debug log is enabled") // only shows if Debug is actually enabled

	//start module from os.Args[1]

	var calendar1 *ics.Calendar
	file1, _ := os.Open("./ical-1.ics")
	calendar1, _ = ics.ParseCalendar(file1)

	var calendar2 *ics.Calendar
	file2, _ := os.Open("./ical-2.ics")
	calendar2, _ = ics.ParseCalendar(file2)

	added, deleted, changed := compare(calendar1, calendar2)

	log.Debug("Added: ", len(added))
	log.Debug("Deleted: ", len(deleted))
	log.Debug("Changed: ", len(changed))

	for _, event := range added {
		fmt.Println("Added:\n\n", event.Serialize())
	}
	for _, event := range deleted {
		fmt.Println("Deleted:\n\n", event.Serialize())
	}
	for _, event := range changed {
		fmt.Println("Changed:\n\n", event.Serialize())
	}
}
