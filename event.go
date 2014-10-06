package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	// organization is the UEM organization that will be checked each time
	organization string
	baseURL      = "http://util.columbiaesc.com/uem"
	respLimit    = 100
)

func init() {
	if organization = os.Getenv("UEM_ORGANIZATION"); organization == "" {
		log.Fatal("'UEM_ORGANIZATION' must be set as an environment variable")
	}
}

// API response from http://util.columbiaesc.com/uem/help
type data struct {
	Data   []Event `json:"data"`
	Status float64 `json:"status"`
}

func getEventData() ([]Event, error) {
	resp, err := http.Get(fmt.Sprintf("%s?group=%s&limit=%d", baseURL, organization, respLimit))
	if err != nil {
		log.Printf("Failed to retrieve data from UEM api => %s", err.Error())
		return []Event{}, fmt.Errorf("ERROR (HTTP) => %s", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// parse data
	var apiResp data
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Printf("Failed to unmarshal data from json resp => %s", err.Error())
		return []Event{}, fmt.Errorf("ERROR (JSON) => %s", err.Error())
	}

	// return event list
	return apiResp.Data, nil
}

type Event struct {
	ID               string `json:"_id"`
	DateStr          string `json:"date_str"`
	EndTime          string `json:"end_time"`
	EndTimeStr       string `json:"end_time_str"`
	Group            string `json:"group"`
	Location         string `json:"location"`
	LocationFull     string `json:"location_full"`
	LocationSpecific string `json:"location_specific"`
	StartTime        string `json:"start_time"`
	StartTimeStr     string `json:"start_time_str"`
	Title            string `json:"title"`
}

func (e Event) Diff(g Event) string {
	var difference string = fmt.Sprintf("\nEvent ID: %s [%s]", e.ID, e.Title)
	if e.DateStr != g.DateStr {
		difference = fmt.Sprintf("%s\nold: %s, new: %s", difference, e.DateStr, g.DateStr)
	}
	if e.EndTime != g.EndTime {
		difference = fmt.Sprintf("%s\nold: %s, new: %s", difference, e.EndTime, g.EndTime)
	}
	if e.EndTimeStr != g.EndTimeStr {
		difference = fmt.Sprintf("%s\nold: %s, new: %s", difference, e.EndTimeStr, g.EndTimeStr)
	}
	if e.Group != g.Group {
		difference = fmt.Sprintf("%s\nold: %s, new: %s", difference, e.Group, g.Group)
	}
	if e.Location != g.Location {
		difference = fmt.Sprintf("%s\nold: %s, new: %s", difference, e.Location, g.Location)
	}
	if e.LocationFull != g.LocationFull {
		difference = fmt.Sprintf("%s\nold: %s, new: %s", difference, e.LocationFull, g.LocationFull)
	}
	if e.LocationSpecific != g.LocationSpecific {
		difference = fmt.Sprintf("%s\nold: %s, new: %s", difference, e.LocationSpecific, g.LocationSpecific)
	}
	if e.StartTime != g.StartTime {
		difference = fmt.Sprintf("%s\nold: %s, new: %s", difference, e.StartTime, g.StartTime)
	}
	if e.StartTimeStr != g.StartTimeStr {
		difference = fmt.Sprintf("%s\nold: %s, new: %s", difference, e.StartTimeStr, g.StartTimeStr)
	}
	if e.Title != g.Title {
		difference = fmt.Sprintf("%s\nold: %s, new: %s", difference, e.Title, g.Title)
	}

	return difference + "\n"
}
