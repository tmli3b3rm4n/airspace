package handler

import (
	"encoding/json"
	"github.com/tmli3b3rm4n/airspace/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) RestrictedAirspace(lat, lon float64) (bool, error) {
	args := m.Called(lat, lon)
	return args.Bool(0), args.Error(1)
}

// TestRestrictedAirspace_ValidCoordinates: testing valid cords.
func TestRestrictedAirspace_ValidCoordinates(t *testing.T) {
	e := echo.New()

	mockRepo := repository.NewMockFlightRestrictionsRepo()
	mockRepo.On("RestrictedAirspace", 32.20, -84.99).Return(true, nil)

	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace/32.20/-84.99", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "success", response["status"])
	assert.Equal(t, map[string]interface{}{"endpoint": "RestrictedAirspace", "value": true}, response["message"])

	mockRepo.AssertExpectations(t)
}

// TestRestrictedAirspace_InvalidLatitude: testing for invalid latitude.
func TestRestrictedAirspace_InvalidLatitude(t *testing.T) {
	e := echo.New()

	mockRepo := repository.NewMockFlightRestrictionsRepo()
	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace/invalid/-84.99", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "Invalid slat", response["message"])
}

// TestRestrictedAirspace_InvalidLongitude: testing for invalid cords.
func TestRestrictedAirspace_InvalidLongitude(t *testing.T) {
	e := echo.New()

	mockRepo := repository.NewMockFlightRestrictionsRepo()
	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace/32.20/invalid", nil)
	rec := httptest.NewRecorder()

	// Perform the request
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "Invalid longitude", response["message"])
}
