package mock

import "github.com/stretchr/testify/mock"

type MockFlightRestrictionsRepo struct {
	mock.Mock
}

// NewMockFlightRestrictionsRepo creates and returns a new instance of MockFlightRestrictionsRepo.
func NewMockFlightRestrictionsRepo() *MockFlightRestrictionsRepo {
	return new(MockFlightRestrictionsRepo)
}

// RestrictedAirspace checks to see if cords are in restrictive airspace.
func (m *MockFlightRestrictionsRepo) RestrictedAirspace(lat, lon float64) (bool, error) {
	args := m.Called(lat, lon)
	return args.Bool(0), args.Error(1)
}
