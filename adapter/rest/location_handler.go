package rest

import (
	"encoding/json"
	"eventbook/core/ports"
	"eventbook/core/services"
	"net/http"
)

type LocationHandler struct {
	locationService ports.LocationHandler
}

func NewLocationHandler(locationService services.LocationService) *LocationHandler {
	return &LocationHandler{locationService: locationService}
}

func (h LocationHandler) GetAllLocations() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := json.Marshal(h.locationService.All())
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
