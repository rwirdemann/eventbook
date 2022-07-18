package sql

import (
	"context"
	"database/sql"
	"eventbook/core/domain"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"time"
)

type EventRepository struct {
	connection *pgx.Conn
}

func NewEventRepository() *EventRepository {
	c, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &EventRepository{connection: c}
}

func (m *EventRepository) All() []domain.Event {
	var events []domain.Event
	rows, _ := m.connection.Query(context.Background(), "select * from events order by date desc")
	for rows.Next() {
		var id int
		var name string
		var location string
		var date time.Time
		var distance sql.NullInt32
		err := rows.Scan(&id, &location, &name, &date, &distance)
		if err != nil {
			panic(err)
		}

		events = append(events, domain.Event{Id: id, Name: name, Location: location, Date: date, Distance: toInt32(distance)})
	}
	return events
}

func toInt32(v sql.NullInt32) int32 {
	if v.Valid {
		return v.Int32
	}
	return 0
}

func (m *EventRepository) CreateOrUpdate(event domain.Event) domain.Event {
	_, err := m.connection.Exec(context.Background(), "insert into events(name, location, date, distance) values($1, $2, $3, $4)", event.Name, event.Location, event.Date, event.Distance)
	if err != nil {
		panic(err)
	}
	return event
}

func (m *EventRepository) Get(id int) domain.Event {
	return domain.Event{}
}

func (m *EventRepository) Delete(id int) {
	_, err := m.connection.Exec(context.Background(), "delete from events where id=$1", id)
	if err != nil {
		panic(err)
	}
}
