package models

import (
	"time"

	"example.com/example/10restAPI/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events []Event = []Event{} // alt of events := []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?,?,?,?,?)` // to avoid sql injection attacks
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)

	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) { //null val for pointer is nil
	query := "SELECT * FROM events WHERE id =?"
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name =?, description =?, location =?, dateTime =?
	WHERE id = ?
	`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

func (event Event) Delete() error {
	query := `
	DELETE FROM events WHERE id = ?
	`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(event.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VAULES (?,?)"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.ID, userId)

	return err
}
