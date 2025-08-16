package models

import (
	"net/http"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserID      int       `json:"userID"`
}

var events []Event

func (e Event) Save() (int, error) {
	// TODO: save event to database
	events = append(events, e)
	return http.StatusCreated, nil
}
