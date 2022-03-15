package services

import (
	"github.com/stretchr/testify/assert"
	"joinapi/adapter"
	"joinapi/core/domain"
	"testing"
)

func TestCreateEvent(t *testing.T) {
	eventRepository := adapter.NewMemoryRepository()
	service := NewEventService(eventRepository)
	event := service.Create(domain.Event{Name: "Heiligenhafen"})
	events := eventRepository.All()
	assert.Equal(t, 1, len(events))
	s := eventRepository.Get(event.Id)
	assert.Equal(t, event.Id, s.Id)
}
