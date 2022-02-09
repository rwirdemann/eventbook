package session

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSession(t *testing.T) {
	sink := NewMemoryDataSink()
	service := NewService(sink)
	session := service.Create(Session{Name: "Heiligenhafen"})
	sessions := sink.All()
	assert.Equal(t, 1, len(sessions))
	s := sink.Get(session.Id)
	assert.Equal(t, session.Id, s.Id)
}
