package services

import (
	"eventbook/adapter/memory"
	"eventbook/core/domain"
)

type RealmService struct {
	realmRepository *memory.RealmRepository
}

func NewRealmService(realmRepository *memory.RealmRepository) RealmService {
	return RealmService{realmRepository: realmRepository}
}

func (s RealmService) All() []domain.Realm {
	return s.realmRepository.All()
}

func (s RealmService) Create(realm domain.Realm) domain.Realm {
	return s.realmRepository.CreateOrUpdate(realm)
}
