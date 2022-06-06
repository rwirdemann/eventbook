package rest

import (
	"encoding/json"
	"eventbook/core/domain"
	"eventbook/core/ports"
	"eventbook/core/services"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EventHandler struct {
	eventService ports.EventHandler
}

func NewEventHandler(eventService services.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

func (h EventHandler) GetAllEvents() http.HandlerFunc {
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

func (h EventHandler) CreateEvent() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if b, err := ioutil.ReadAll(request.Body); err == nil {
			if len(b) == 0 {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			var event domain.Event
			_ = json.Unmarshal(b, &event)
			e := h.eventService.Create(event)
			url := request.URL.String()
			writer.Header().Set("Location", fmt.Sprintf("%s/%d", url, e.Id))
			writer.WriteHeader(http.StatusCreated)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
