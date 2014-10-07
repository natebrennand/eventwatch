package main

import (
	"fmt"
	_ "github.com/natebrennand/admin" // add healthcheck endpoint
	"log"
	"time"
)

var (
	// checkInterval is the number of seconds between each check on UEM
	checkInterval = time.Minute * 60
)

// eventLog holds the cache of the UEM events
type eventLog struct {
	events map[string]Event
}

// check compares the current cached events against the most recently returned set from UEM
func (e eventLog) check(newEvents []Event) error {
	var errMsg = ""
	for _, event := range newEvents {
		old, exists := e.events[event.ID]

		// add if new
		if !exists {
			e.events[event.ID] = event
			continue
		}

		// check if changed
		if event != old {
			errMsg += old.Diff(event)
		}
	}

	// return error if things are changed
	if errMsg != "" {
		return fmt.Errorf("Events Changed:\n%s", errMsg)
	}
	return nil
}

func main() {
	eventCache := eventLog{
		events: make(map[string]Event),
	}

	// start a useless web frontend to appease
	go helloworld()

	for {
		log.Println("Starting to check for changes")
		newEvents, err := getEventData()
		if err != nil {
			log.Println(err.Error())
		}

		if err = eventCache.check(newEvents); err != nil {
			log.Println("FOUND CHANGES, emailing user")
			log.Println(err.Error())
			notify(err.Error())
		} else {
			log.Println("Checked for changes and found none")
		}

		time.Sleep(checkInterval)
	}
}
