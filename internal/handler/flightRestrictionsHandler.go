package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/airspace/internal/repository/flightRestrictions"
	"github.com/tmli3b3rm4n/airspace/pkg/cords"
	"github.com/tmli3b3rm4n/airspace/pkg/response"
	"net/http"
)

// FlightRestrictionsHandler handles flight restriction related requests
type FlightRestrictionsHandler struct {
	repo flightRestrictions.IFlightRestrictions
}

// NewFlightRestrictionsHandler creates a new handler
func NewFlightRestrictionsHandler(repo flightRestrictions.IFlightRestrictions) *FlightRestrictionsHandler {
	return &FlightRestrictionsHandler{repo}
}

// RestrictedAirspace checks if the provided coordinates are in restricted airspace
func (f *FlightRestrictionsHandler) RestrictedAirspace(c echo.Context) error {
	lat, lon, err := cords.GetLatLon(c)

	isRestricted, err := f.repo.RestrictedAirspace(lat, lon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{
			Status: "error",
			Error:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response{
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
