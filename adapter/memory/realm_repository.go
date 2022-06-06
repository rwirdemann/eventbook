package memory

import "eventbook/core/domain"

type RealmRepository struct {
	realms map[int]domain.Realm
	id     int
}

func NewRealmRepository() *RealmRepository {
	return &RealmRepository{realms: make(map[int]domain.Realm), id: 1}
}

func (m *RealmRepository) All() []domain.Realm {
	var sessions []domain.Realm
	for _, v := range m.realms {
		sessions = append(sessions, v)
	}
	return sessions
}

func (m *RealmRepository) CreateOrUpdate(realm domain.Realm) domain.Realm {
	maxID := 1
	for k, _ := range m.realms {
		if k > maxID {
			maxID = k
		}
	}
	realm.Id = maxID + 1
	m.realms[realm.Id] = realm
	m.id++
	return realm
}

func (m *RealmRepository) Get(id int) domain.Realm {
	return m.realms[id]
}
