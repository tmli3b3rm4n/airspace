package handler

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	mock2 "github.com/tmli3b3rm4n/airspace/internal/repository/flightRestrictions/mock"
	"github.com/tmli3b3rm4n/airspace/pkg/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestRestrictedAirspace_ValidCoordinates: testing valid cords.
func TestRestrictedAirspace_ValidCoordinates(t *testing.T) {
	e := echo.New()

	mockRepo := mock2.NewMockFlightRestrictionsRepo()
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

// TestRestrictedAirspace_InvalidLatitude : testing invalid cords
func TestRestrictedAirspace_InvalidLatitude(t *testing.T) {
	mockRepo := new(mock2.MockFlightRestrictionsRepo) // Create the mock repository
	handler := NewFlightRestrictionsHandler(mockRepo)

	// Define the behavior of the mock
	mockRepo.On("RestrictedAirspace", mock.AnythingOfType("float64"), mock.AnythingOfType("float64")).
		Return(false, errors.New("repository error")) // Return a valid bool and an error

	// Setup the request and recorder
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace?lat=0&long=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler
	err := handler.RestrictedAirspace(c)
	assert.NoError(t, err)

	// Validate the response
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	var response response.Response
	err = json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "error", response.Status)
	assert.Equal(t, "repository error", response.Error)

	// Verify the mock's expectations
	mockRepo.AssertExpectations(t)
}
