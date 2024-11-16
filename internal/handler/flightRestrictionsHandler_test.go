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

// Mocking the repository (f.repo)
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) RestrictedAirspace(lat, lon float64) (bool, error) {
	args := m.Called(lat, lon)
	return args.Bool(0), args.Error(1)
}
func TestRestrictedAirspace_ValidCoordinates(t *testing.T) {
	// Setup Echo instance and mock repository
	e := echo.New()

	// Use the newMockFlightRestrictionsRepo function to create a new mock
	mockRepo := repository.NewMockFlightRestrictionsRepo()
	mockRepo.On("RestrictedAirspace", 32.20, -84.99).Return(true, nil)

	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	// Prepare the test request with valid coordinates
	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace/32.20/-84.99", nil)
	rec := httptest.NewRecorder()

	// Perform the request
	e.ServeHTTP(rec, req)

	// Check if the response code is OK (200)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Check if the response body is as expected
	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	// Check if the response message is as expected
	assert.Equal(t, "success", response["status"])
	assert.Equal(t, map[string]interface{}{"endpoint": "RestrictedAirspace", "value": true}, response["message"])

	// Verify that the repository method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestRestrictedAirspace_InvalidLatitude(t *testing.T) {
	// Setup Echo instance and mock repository
	e := echo.New()

	// Use the newMockFlightRestrictionsRepo function to create a new mock
	mockRepo := repository.NewMockFlightRestrictionsRepo()
	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	// Prepare the test request with an invalid latitude
	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace/invalid/-84.99", nil)
	rec := httptest.NewRecorder()

	// Perform the request
	e.ServeHTTP(rec, req)

	// Check if the response code is BadRequest (400)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Check if the response body is as expected
	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	// Check the specific error message
	assert.Equal(t, "Invalid slat", response["message"])
}

func TestRestrictedAirspace_InvalidLongitude(t *testing.T) {
	// Setup Echo instance and mock repository
	e := echo.New()

	// Use the newMockFlightRestrictionsRepo function to create a new mock
	mockRepo := repository.NewMockFlightRestrictionsRepo()
	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	// Prepare the test request with an invalid longitude
	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace/32.20/invalid", nil)
	rec := httptest.NewRecorder()

	// Perform the request
	e.ServeHTTP(rec, req)

	// Check if the response code is BadRequest (400)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Check if the response body is as expected
	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	// Check the specific error message
	assert.Equal(t, "Invalid longitude", response["message"])
}
