package ports

import "eventbook/core/domain"

type LocationHandler interface {
	All() []domain.Location
}
