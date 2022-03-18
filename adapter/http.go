package adapter

import (
	"encoding/json"
	"eventbook/core/services"
	"net/http"
)

type HTTPHandler struct {
	eventService services.EventService
	realmService services.RealmService
}

func NewHTTPHandler(eventService services.EventService, realmService services.RealmService) *HTTPHandler {
	return &HTTPHandler{eventService: eventService, realmService: realmService}
}

func (h HTTPHandler) GetAllEvents() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := json.Marshal(h.eventService.All())
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(bytes)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (h HTTPHandler) GetAllRealms() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := json.Marshal(h.realmService.All())
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(bytes)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
