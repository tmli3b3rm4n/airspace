package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/airspace/internal/repository"
	"log"
	"net/http"
	"strconv"
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
