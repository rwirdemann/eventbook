package services

import (
	"eventbook/core/domain"
)

type RealmService struct {
}

func NewRealmService() RealmService {
	return RealmService{}
}

func (s RealmService) All() []domain.Realm {
	return []domain.Realm{{Name: "Wingbuddies"}}
}
