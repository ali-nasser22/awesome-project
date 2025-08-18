package models

import (
	"awesomeProject/db"
	"net/http"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime"`
	UserID      string    `json:"userID" binding:"required"`
}

var events []Event

func (e Event) Save() (int, error) {
	query := "INSERT INTO events (name,description,dateTime, location,userId) VALUES (?, ?,?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return http.StatusBadRequest, err
	}
	defer stmt.Close()
	e.DateTime = time.Now()
	result, err := stmt.Exec(e.Name, e.Description, e.DateTime, e.Location, e.UserID)
	if err != nil {
		return http.StatusBadRequest, err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return http.StatusCreated, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.UserID,
			&event.DateTime,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil

}

func GetEventById(id int64) (Event, error) {
	query := "SELECT * FROM events WHERE id=?"

	result := db.DB.QueryRow(query, id)
	var event Event
	err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.UserID, &event.DateTime)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}
