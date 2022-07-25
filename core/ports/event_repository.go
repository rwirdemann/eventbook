package ports

import (
	"eventbook/core/domain"
)

type EventRepository interface {
	Create(session domain.Event) domain.Event
	Update(id int, event domain.Event) domain.Event
	All() []domain.Event
	Get(id int) domain.Event
	Delete(id int)
}
