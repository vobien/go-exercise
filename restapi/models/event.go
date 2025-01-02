package models

import "time"

type Event struct {
	Id          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string
	DateTime    time.Time
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetEvents() []Event {
	return events
}
