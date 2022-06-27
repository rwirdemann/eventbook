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
