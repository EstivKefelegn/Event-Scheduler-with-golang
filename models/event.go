package models

import (
	"time"

	"Eventplanning.go/Api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"Name" binding:"required"`
	Description string    `json:"Description" binding:"required"`
	Location    string    `json:"Location" binding:"required"`
	Datetime    time.Time `json:"Datetime" binding:"required"`
	UserID      int64       `json:"UserID"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, datetime, user_id)
		VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserID)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE ID = ?`
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, datetime = ?
	WHERE id = ? 
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.Datetime, event.UserID)

	return err
}

func (event Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.ID)

	return err

}
