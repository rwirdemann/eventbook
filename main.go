package main

import (
	"eventbook/adapter"
	"eventbook/core/domain"
	"eventbook/core/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	eventService := services.NewEventService(adapter.NewMemoryRepository())
	realmService := services.NewRealmService()
	eventService.Create(domain.Event{Name: "Wingding Heiligenhafen"})
	httpAdapter := adapter.NewHTTPHandler(eventService, realmService)
	router := mux.NewRouter()
	router.HandleFunc("/admin/realms", adapter.JWTAuth(httpAdapter.GetAllRealms()))
	router.HandleFunc("/events", httpAdapter.GetAllEvents())
	log.Fatal(http.ListenAndServe(":8000", router))
}
