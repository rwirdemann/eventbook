package rest

import (
	"encoding/json"
	"eventbook/core/domain"
	"eventbook/core/ports"
	"eventbook/core/services"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
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

func (h LocationHandler) DeleteLocation() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id, _ := strconv.Atoi(params["id"])
		h.locationService.Delete(id)
		writer.WriteHeader(http.StatusOK)
	}
}

func (h LocationHandler) CreateLocation() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if b, err := ioutil.ReadAll(request.Body); err == nil {
			if len(b) == 0 {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			var location domain.Location
			_ = json.Unmarshal(b, &location)

			l, exists := h.locationService.FindByName(location.Name)
			if exists {
				url := request.URL.String()
				writer.Header().Set("Location", fmt.Sprintf("%s/%d", url, l.Id))
				writer.WriteHeader(http.StatusOK)
			} else {
				e := h.locationService.Create(location)
				url := request.URL.String()
				writer.Header().Set("Location", fmt.Sprintf("%s/%d", url, e.Id))
				writer.WriteHeader(http.StatusCreated)
			}
		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
