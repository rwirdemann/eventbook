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
	return s.eventRepository.CreateOrUpdate(event)
}

func (s EventService) All() []domain.Event {
	return s.eventRepository.All()
}
