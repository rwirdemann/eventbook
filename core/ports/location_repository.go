package ports

import (
	"eventbook/core/domain"
)

type LocationRepository interface {
	All() []domain.Location
}
