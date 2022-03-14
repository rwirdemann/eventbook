package event

type Service struct {
	eventRepository Repository
}

func NewService(eventRepository Repository) Service {
	return Service{eventRepository: eventRepository}
}

func (s Service) Create(session Event) Event {
	return s.eventRepository.CreateOrUpdate(session)
}

func (s Service) All() []Event {
	return s.eventRepository.All()
}
