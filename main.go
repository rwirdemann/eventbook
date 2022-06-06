package main

import (
	"eventbook/adapter/memory"
	"eventbook/adapter/rest"
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

	eventService := services.NewEventService(memory.NewEventRepository())

	realmService := services.NewRealmService(memory.NewRealmRepository())
	realmService.Create(domain.Realm{Name: "Wingbuddies"})
	realmService.Create(domain.Realm{Name: "Bikebuddies"})

	eventAdapter := rest.NewEventHandler(eventService)
	realmAdapter := rest.NewRealmHandler(realmService)
	router := mux.NewRouter()
	router.HandleFunc("/admin/realms", rest.JWTAuth(realmAdapter.GetAllRealms())).Methods("GET")
	router.HandleFunc("/admin/realms", rest.JWTAuth(realmAdapter.CreateRealm())).Methods("POST")
	router.HandleFunc("/events", eventAdapter.GetAllEvents()).Methods("GET")
	router.HandleFunc("/events", eventAdapter.CreateEvent()).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
