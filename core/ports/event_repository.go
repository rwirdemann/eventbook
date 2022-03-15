package ports

import (
	"joinapi/core/domain"
)

type EventRepository interface {
	CreateOrUpdate(session domain.Event) domain.Event
	All() []domain.Event
	Get(id int) domain.Event
}
