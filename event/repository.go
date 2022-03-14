package event

type Repository interface {
	CreateOrUpdate(session Event) Event
	All() []Event
	Get(id int) Event
}

type MemoryRepository struct {
	sessions map[int]Event
	id       int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{sessions: make(map[int]Event), id: 1}
}

func (m *MemoryRepository) All() []Event {
	var sessions []Event
	for _, v := range m.sessions {
		sessions = append(sessions, v)
	}
	return sessions
}

func (m *MemoryRepository) CreateOrUpdate(session Event) Event {
	session.Id = m.id
	m.sessions[m.id] = session
	m.id++
	return session
}

func (m *MemoryRepository) Get(id int) Event {
	return m.sessions[id]
}
