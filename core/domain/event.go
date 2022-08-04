package domain

import "time"

type Event struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Distance float64   `json:"distance"`
	Date     time.Time `json:"date"`
	MaxSpeed float64   `json:"maxspeed"`
	Duration float64   `json:"duration"`
}
