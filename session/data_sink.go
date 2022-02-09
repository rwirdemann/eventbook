package session

type DataSink interface {
	CreateOrUpdate(session Session) Session
	All() []Session
	Get(id int) Session
}

type MemoryDataSink struct {
	sessions map[int]Session
	id       int
}

func NewMemoryDataSink() *MemoryDataSink {
	return &MemoryDataSink{sessions: make(map[int]Session), id: 1}
}

func (m *MemoryDataSink) All() []Session {
	var sessions []Session
	for _, v := range m.sessions {
		sessions = append(sessions, v)
	}
	return sessions
}

func (m *MemoryDataSink) CreateOrUpdate(session Session) Session {
	session.Id = m.id
	m.sessions[m.id] = session
	m.id++
	return session
}

func (m *MemoryDataSink) Get(id int) Session {
	return m.sessions[id]
}
