package adapter

import (
	"encoding/json"
	"eventbook/core/services"
	"net/http"
)

type HTTPHandler struct {
	eventService services.EventService
}

func NewHTTPHandler(eventService services.EventService) *HTTPHandler {
	return &HTTPHandler{eventService: eventService}
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
