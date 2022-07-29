package sql

import (
	"context"
	"eventbook/core/domain"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type LocationRepository struct {
	connection *pgx.Conn
}

func NewLocationRepository() *LocationRepository {
	c, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &LocationRepository{connection: c}
}

func (m *LocationRepository) All() []domain.Location {
	var locations []domain.Location
	rows, _ := m.connection.Query(context.Background(), "select * from locations order by name")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		locations = append(locations, domain.Location{Id: id, Name: name})
	}
	return locations
}

func (m *LocationRepository) FindByName(name string) (domain.Location, bool) {
	rows, _ := m.connection.Query(context.Background(), "select * from locations where name = $1", name)
	if rows.Next() {
		var l domain.Location
		err := rows.Scan(&l.Id, &l.Name)
		if err != nil {
			panic(err)
		}
		return l, true
	}
	return domain.Location{}, false
}

func (m *LocationRepository) Create(location domain.Location) {
	_, err := m.connection.Exec(context.Background(), "insert into locations(name) values($1)", location.Name)
	if err != nil {
		panic(err)
	}
}
