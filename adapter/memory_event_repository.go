package adapter

import "eventbook/core/domain"

type MemoryEventRepository struct {
	sessions map[int]domain.Event
	id       int
}

func NewMemoryRepository() *MemoryEventRepository {
	return &MemoryEventRepository{sessions: make(map[int]domain.Event), id: 1}
}

func (m *MemoryEventRepository) All() []domain.Event {
	var sessions []domain.Event
	for _, v := range m.sessions {
		sessions = append(sessions, v)
	}
	return sessions
}

func (m *MemoryEventRepository) CreateOrUpdate(session domain.Event) domain.Event {
	session.Id = m.id
	m.sessions[m.id] = session
	m.id++
	return session
}

func (m *MemoryEventRepository) Get(id int) domain.Event {
	return m.sessions[id]
}
