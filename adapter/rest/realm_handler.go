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

type RealmHandler struct {
	realmService ports.RealmHandler
}

func NewRealmHandler(realmService services.RealmService) *RealmHandler {
	return &RealmHandler{realmService: realmService}
}

func (h RealmHandler) GetAllRealms() func(http.ResponseWriter, *http.Request) {
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

func (h RealmHandler) CreateRealm() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		var realm domain.Realm
		_ = json.Unmarshal(body, &realm)

		created := h.realmService.Create(realm)
		b, err := json.Marshal(created)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		location := fmt.Sprintf("%s/%d", request.URL.String(), created.Id)
		writer.Header().Set("Location", location)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		_, _ = writer.Write(b)
	}
}
