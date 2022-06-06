package ports

import "eventbook/core/domain"

type EventHandler interface {
	Create(session domain.Event) domain.Event
	All() []domain.Event
}
