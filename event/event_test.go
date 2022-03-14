package event

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEvent(t *testing.T) {
	eventRepository := NewMemoryRepository()
	service := NewService(eventRepository)
	event := service.Create(Event{Name: "Heiligenhafen"})
	events := eventRepository.All()
	assert.Equal(t, 1, len(events))
	s := eventRepository.Get(event.Id)
	assert.Equal(t, event.Id, s.Id)
}
