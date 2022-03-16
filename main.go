package main

import (
	"eventbook/adapter"
	"eventbook/core/domain"
	"eventbook/core/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	service := services.NewEventService(adapter.NewMemoryRepository())
	service.Create(domain.Event{Name: "Wingding Heiligenhafen"})
	httpAdapter := adapter.NewHTTPHandler(service)
	router := mux.NewRouter()
	router.HandleFunc("/events", httpAdapter.GetAllEvents())
	log.Fatal(http.ListenAndServe(":8000", router))
}
