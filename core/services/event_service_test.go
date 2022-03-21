package services

import (
	"eventbook/adapter/memory"
	"eventbook/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEvent(t *testing.T) {
	eventRepository := memory.NewEventRepository()
	service := NewEventService(eventRepository)
	event := service.Create(domain.Event{Name: "Heiligenhafen"})
	events := eventRepository.All()
	assert.Equal(t, 1, len(events))
	s := eventRepository.Get(event.Id)
	assert.Equal(t, event.Id, s.Id)
}
