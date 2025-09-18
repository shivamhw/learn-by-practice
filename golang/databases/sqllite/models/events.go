package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Events struct {
	db *sql.DB
}

type EventModel struct {
	Id          int64
	Owner_id    int64
	Name        string
	Description string
	Date        time.Time
}

func NewEventDb(db *sql.DB) *Events {
	return &Events{
		db: db,
	}
}

func (e *Events) Insert(event EventModel) (*EventModel, error) {
	res, err := e.db.ExecContext(context.Background(), "INSERT INTO events (owner_id, name, description, date) VALUES (?, ?, ?, ?)", event.Owner_id, event.Name, event.Description, event.Date)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	return &EventModel{
		Id: id,
		Owner_id: event.Owner_id,
		Name: event.Name,
		Description: event.Description,
		Date: event.Date,
	}, nil
}

func (e *Events) List() (events []*EventModel, err error) {
	rows, err := e.db.QueryContext(context.Background(), "SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		t := &EventModel{}
		rows.Scan(&t.Id, &t.Owner_id, &t.Name, &t.Description, &t.Date)
		events = append(events, t)
	}
	return
}

func (e *Events) Get(id int64) (event EventModel, err error) {
	err = e.db.QueryRowContext(context.Background(), "SELECT * FROM events where id = ?", id).Scan(&event.Id, &event.Owner_id, &event.Name, &event.Description, &event.Date)
	if err == sql.ErrNoRows {
		fmt.Printf("no rows found for %d", id)
		return
	}
	if err != nil {
		return 
	}
	return 
}