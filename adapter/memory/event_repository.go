package memory

import "eventbook/core/domain"

type EventRepository struct {
	sessions map[int]domain.Event
	id       int
}

func NewEventRepository() *EventRepository {
	return &EventRepository{sessions: make(map[int]domain.Event), id: 1}
}

func (m *EventRepository) All() []domain.Event {
	var sessions []domain.Event
	for _, v := range m.sessions {
		sessions = append(sessions, v)
	}
	return sessions
}

func (m *EventRepository) CreateOrUpdate(session domain.Event) domain.Event {
	session.Id = m.id
	m.sessions[m.id] = session
	m.id++
	return session
}

func (m *EventRepository) Get(id int) domain.Event {
	return m.sessions[id]
}