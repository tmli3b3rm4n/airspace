package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/airspace/internal/repository/flightRestrictions"
	"github.com/tmli3b3rm4n/airspace/pkg/parse"
	"github.com/tmli3b3rm4n/airspace/pkg/response"
	"github.com/tmli3b3rm4n/airspace/pkg/util"
)

// FlightRestrictionsHandler handles flight restriction related requests
type FlightRestrictionsHandler struct {
	repo flightRestrictions.IFlightRestrictions
}

// NewFlightRestrictionsHandler creates a new handler
func NewFlightRestrictionsHandler(repo flightRestrictions.IFlightRestrictions) *FlightRestrictionsHandler {
	return &FlightRestrictionsHandler{repo}
}

func (f *FlightRestrictionsHandler) RestrictedAirspace(c echo.Context) error {
	// Parse and validate lat/lon
	lat, lon, err := parse.ParseLatLon(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status: "invalid latitude or longitude.",
			Error:  "failed to parse latitude or longitude.",
		})
	}

	// Validate latitude
	if !util.IsValidLatitude(lat) {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status: "invalid latitude.",
			Error:  "latitude out of range.",
		})
	}

	// Validate longitude
	if !util.IsValidLongitude(lon) {
		return c.JSON(http.StatusBadRequest, response.Response{
			Status: "invalid longitude.",
			Error:  "longitude out of range.",
		})
	}

	// Check repository for restricted airspace
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
