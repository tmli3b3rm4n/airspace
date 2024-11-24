package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mock2 "github.com/tmli3b3rm4n/airspace/internal/repository/flightRestrictions/mock"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRestrictedAirspace(t *testing.T) {
	e := echo.New()
	mockRepo := new(mock2.MockFlightRestrictionsRepo)
	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	tests := []struct {
		name           string
		url            string
		mockReturn     interface{}
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Valid Coordinates",
			url:            "/restricted-airspace/32.20/-84.99",
			mockReturn:     true,
			expectedStatus: http.StatusOK,
			expectedError:  "",
		},
		{
			name:           "Invalid Latitude",
			url:            "/restricted-airspace/invalid/0",
			mockReturn:     nil, // Will not reach the repo
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid latitude or longitude.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockReturn != nil {
				mockRepo.On("RestrictedAirspace", mock.Anything, mock.Anything).
					Return(tc.mockReturn, nil).Once()
			}

			req := httptest.NewRequest(http.MethodGet, tc.url, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)

			if tc.expectedError != "" {
				var response map[string]interface{}
				err := json.NewDecoder(rec.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedError, response["status"])
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
func TestRestrictedAirspace_Invalid(t *testing.T) {
	e := echo.New()
	mockRepo := new(mock2.MockFlightRestrictionsRepo) // Mock repository
	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace/invalid/0", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "invalid latitude or longitude.", response["status"])
	assert.Equal(t, "failed to parse latitude or longitude.", response["error"])
}

func TestRestrictedAirspace_InvalidLatitude(t *testing.T) {
	e := echo.New()
	mockRepo := new(mock2.MockFlightRestrictionsRepo) // Mock repository
	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace/-3000.333/5000.333", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "invalid latitude.", response["status"])
	assert.Equal(t, "latitude out of range.", response["error"])
}

func TestRestrictedAirspace_InvalidLongitude(t *testing.T) {
	e := echo.New()
	mockRepo := new(mock2.MockFlightRestrictionsRepo) // Mock repository
	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace/32.20/5000.333", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "invalid longitude.", response["status"])
	assert.Equal(t, "longitude out of range.", response["error"])
}
