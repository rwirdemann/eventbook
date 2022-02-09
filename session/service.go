package session

type Service struct {
	dataSink DataSink
}

func NewService(port DataSink) Service {
	return Service{dataSink: port}
}

func (s Service) Create(session Session) Session {
	return s.dataSink.CreateOrUpdate(session)
}
