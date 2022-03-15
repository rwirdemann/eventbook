package main

import (
	"joinapi/adapter"
	"joinapi/core/domain"
	"joinapi/core/services"
)

func main() {
	service := services.NewEventService(adapter.NewMemoryRepository())
	service.Create(domain.Event{Name: "Heilgenhafen"})
}
