package main

import "joinapi/event"

func main() {
	service := event.NewService(&event.MemoryRepository{})
	service.Create(event.Event{Name: "Heilgenhafen"})
}
