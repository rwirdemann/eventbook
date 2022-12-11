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
	defer rows.Close()
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
	defer rows.Close()
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

func (m *LocationRepository) Create(location domain.Location) domain.Location {
	sqlStatement := "insert into locations(name) values($1) returning id"
	id := 0
	err := m.connection.QueryRow(context.Background(),
		sqlStatement, location.Name).Scan(&id)
	if err != nil {
		panic(err)
	}
	location.Id = id
	return location
}

func (m *LocationRepository) Delete(id int) {
	_, err := m.connection.Exec(context.Background(), "delete from locations where id=$1", id)
	if err != nil {
		panic(err)
	}
}
