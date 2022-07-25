package services

import (
	"eventbook/core/domain"
	"eventbook/core/ports"
)

type EventService struct {
	eventRepository ports.EventRepository
}

func NewEventService(eventRepository ports.EventRepository) EventService {
	return EventService{eventRepository: eventRepository}
}

func (s EventService) Create(event domain.Event) domain.Event {
	return s.eventRepository.Create(event)
}

func (s EventService) All() []domain.Event {
	events := s.eventRepository.All()
	if events == nil {
		return []domain.Event{}
	}
	return events
}

func (s EventService) Delete(id int) {
	s.eventRepository.Delete(id)
}
