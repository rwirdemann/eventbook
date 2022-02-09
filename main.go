package main

import "joinapi/session"

func main() {
	service := session.NewService(session.MemoryDataSink{})
	service.Create(session.Session{Name: "Heilgenhafen"})
}
