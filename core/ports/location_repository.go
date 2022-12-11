package ports

import (
	"eventbook/core/domain"
)

type LocationRepository interface {
	All() []domain.Location
	FindByName(name string) (domain.Location, bool)
	Create(location domain.Location) domain.Location
	Delete(id int)
}
