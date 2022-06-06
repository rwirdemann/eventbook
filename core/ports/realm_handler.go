package ports

import "eventbook/core/domain"

type RealmHandler interface {
	Create(realm domain.Realm) domain.Realm
	All() []domain.Realm
}
