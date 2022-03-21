package ports

import (
	"eventbook/core/domain"
)

type RealmRepository interface {
	CreateOrUpdate(realm domain.Realm) domain.Event
	All() []domain.Realm
	Get(id int) domain.Realm
}
