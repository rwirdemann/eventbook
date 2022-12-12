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
	rows, _ := m.connection.Query(context.Background(),
		"select events.id, events.name, events.date, events.distance, events.maxspeed, events.duration, locations.id, locations.name from events join locations on location_id = locations.id order by events.date desc")
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var location string
		var date time.Time
		var distance sql.NullFloat64
		var maxSpeed sql.NullFloat64
		var duration sql.NullFloat64
		var locationId sql.NullInt32

		err := rows.Scan(&id, &name, &date, &distance, &maxSpeed, &duration, &locationId, &location)
		if err != nil {
			panic(err)
		}

		events = append(events, domain.Event{Id: id, Name: name, Date: date, Distance: toFloat64(distance),
			MaxSpeed: toFloat64(maxSpeed), Duration: toFloat64(duration), LocationId: toInt32(locationId), Location: location})
	}
	return events
}

func toFloat64(v sql.NullFloat64) float64 {
	if v.Valid {
		return v.Float64
	}
	return 0.0
}

func toInt32(v sql.NullInt32) int32 {
	if v.Valid {
		return v.Int32
	}
	return 0
}

func (m *EventRepository) Create(event domain.Event) domain.Event {
	sqlStatement := `
		INSERT INTO events (name, location, date, distance, maxspeed, duration, location_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	id := 0
	err := m.connection.QueryRow(context.Background(),
		sqlStatement, event.Name, event.Location, event.Date, event.Distance, event.MaxSpeed, event.Duration, event.LocationId).Scan(&id)
	if err != nil {
		panic(err)
	}
	event.Id = id
	return event
}

func (m *EventRepository) Update(id int, event domain.Event) domain.Event {
	_, err := m.connection.Exec(context.Background(), "update events set name=$1, location=$2, date=$3, distance=$4, maxspeed=$5, duration=$6, location_id=$7 where id = $8",
		event.Name, event.Location, event.Date, event.Distance, event.MaxSpeed, event.Duration, event.LocationId, id)
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
