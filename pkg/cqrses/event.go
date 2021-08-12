package cqrses

import (
	"github.com/google/uuid"
	"time"
)

type Event interface {
	GetID() string

	GetName() string

	GetTimestamp() time.Time

	GetData() string
}

type EventData struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
}

func (evt EventData) GetID() string {
	return evt.ID
}

func (evt EventData) GetName() string {
	return evt.Name
}

func (evt EventData) GetTimestamp() time.Time {
	return evt.Timestamp
}

func (evt EventData) GetData() string {
	return evt.Data
}

var noopTime = time.Time{}

func NewEvent(evt *EventData) Event {
	if evt.Timestamp == noopTime {
		evt.Timestamp = time.Now()
	}
	if evt.ID == "" {
		evt.ID = uuid.NewString()
	}
	return evt
}
