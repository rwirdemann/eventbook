package services

import (
	"eventbook/core/domain"
	"eventbook/core/ports"
)

type LocationService struct {
	locationRepository ports.LocationRepository
}

func NewLocationService(locationRepository ports.LocationRepository) LocationService {
	return LocationService{locationRepository: locationRepository}
}

func (s LocationService) All() []domain.Location {
	locations := s.locationRepository.All()
	if locations == nil {
		return []domain.Location{}
	}
	return locations
}

func (s LocationService) Delete(id int) {
	s.locationRepository.Delete(id)
}

func (s LocationService) Create(location domain.Location) domain.Location {
	e := s.locationRepository.Create(location)
	return e
}

func (s LocationService) FindByName(name string) (domain.Location, bool) {
	return s.locationRepository.FindByName(name)
}
