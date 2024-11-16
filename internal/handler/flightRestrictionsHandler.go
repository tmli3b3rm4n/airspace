package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tmli3b3rm4n/airspace/internal/repository"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// FlightRestrictionsHandler handles flight restriction related requests
type FlightRestrictionsHandler struct {
	repo repository.IFlightRestrictions
}

// NewFlightRestrictionsHandler creates a new handler
func NewFlightRestrictionsHandler(repo repository.IFlightRestrictions) *FlightRestrictionsHandler {
	return &FlightRestrictionsHandler{repo}
}

// Response represents the API response structure
type Response struct {
	Status  string `json:"status"`
	Message struct {
		Endpoint string `json:"endpoint"`
		Value    bool   `json:"value"`
	} `json:"message"`
	Error string `json:"error,omitempty"` // Optional error field
}

// TestRestrictedAirspace_ErrorFromRepository :
func TestRestrictedAirspace_ErrorFromRepository(t *testing.T) {
	e := echo.New()

	mockRepo := new(repository.MockFlightRestrictionsRepo)
	mockRepo.On("RestrictedAirspace", 32.20, -84.99).Return(false, fmt.Errorf("database error"))

	handler := &FlightRestrictionsHandler{repo: mockRepo}
	e.GET("/restricted-airspace/:lat/:lon", handler.RestrictedAirspace)

	req := httptest.NewRequest(http.MethodGet, "/restricted-airspace/32.20/-84.99", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, "database error", response["error"])
}

// RestrictedAirspace checks if the provided coordinates are in restricted airspace
func (f *FlightRestrictionsHandler) RestrictedAirspace(c echo.Context) error {
	slat := c.Param("lat")
	slon := c.Param("lon")

	lat, err := strconv.ParseFloat(slat, 64)
	log.Printf("slat : before parse %v,  %v", c.Param("lat"), slat)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid slat"})
	}

	lon, err := strconv.ParseFloat(slon, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid longitude"})
	}

	isRestricted := false
	isRestricted, err = f.repo.RestrictedAirspace(lat, lon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status: "error",
			Error:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status: "success",
		Message: struct {
			Endpoint string `json:"endpoint"`
			Value    bool   `json:"value"`
		}{
			Endpoint: "RestrictedAirspace",
			Value:    isRestricted,
		},
	})
}
