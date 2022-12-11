package ports

import "eventbook/core/domain"

type LocationHandler interface {
	All() []domain.Location
	Delete(id int)
	Create(location domain.Location) domain.Location
	FindByName(name string) (domain.Location, bool)
}
