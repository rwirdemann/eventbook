package domain

import "time"

type Event struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}
