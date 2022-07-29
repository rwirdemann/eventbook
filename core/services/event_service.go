package services

import (
	"eventbook/core/domain"
	"eventbook/core/ports"
)

type EventService struct {
	eventRepository    ports.EventRepository
	locationRepository ports.LocationRepository
}

func NewEventService(eventRepository ports.EventRepository, locationRepository ports.LocationRepository) EventService {
	return EventService{eventRepository: eventRepository, locationRepository: locationRepository}
}

func (s EventService) Create(event domain.Event) domain.Event {
	e := s.eventRepository.Create(event)
	_, exists := s.locationRepository.FindByName(e.Location)
	if !exists {
		s.locationRepository.Create(domain.Location{
			Name: e.Location,
		})
	}
	return e
}

func (s EventService) Update(id int, event domain.Event) domain.Event {
	return s.eventRepository.Update(id, event)
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
