package services

import (
	"joinapi/core/domain"
	"joinapi/core/ports"
)

type EventService struct {
	eventRepository ports.EventRepository
}

func NewEventService(eventRepository ports.EventRepository) EventService {
	return EventService{eventRepository: eventRepository}
}

func (s EventService) Create(session domain.Event) domain.Event {
	return s.eventRepository.CreateOrUpdate(session)
}

func (s EventService) All() []domain.Event {
	return s.eventRepository.All()
}
