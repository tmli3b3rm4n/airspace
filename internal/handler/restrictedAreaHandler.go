package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/airspace/internal/repository"
	"github.com/tmli3b3rm4n/airspace/pkg/formulas"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RestrictedAreaHandler struct {
	Repo repository.IRestrictedAreaRepository
}

func NewRestrictedAreaHandler(repo repository.IRestrictedAreaRepository) *RestrictedAreaHandler {
	return &RestrictedAreaHandler{
		Repo: repo,
	}
}

func (h *RestrictedAreaHandler) GetByArea(c echo.Context) error {
	areaID, err := strconv.ParseInt(c.Param("areaID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("error: %v", err))
	}

	area, err := h.Repo.FindByID(context.Background(), uint(areaID)) // Replace StrToUint with actual conversion logic
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Area not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the area data in the response
	return c.JSON(http.StatusOK, area)
}

func (h *RestrictedAreaHandler) IsPointInRestrictedArea(c echo.Context) error {
	lat := c.Param("lat")
	lon := c.Param("lon")

	latFloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid latitude")
	}

	lonFloat, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid longitude")
	}

	areas, err := h.Repo.FindAll(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch restricted areas")
	}

	for _, area := range areas {
		distance := formulas.Haversine(latFloat, lonFloat, area.Latitude, area.Longitude)
		if distance <= area.Radius {
			return c.JSON(http.StatusOK, map[string]interface{}{"in_restricted_area": true})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"in_restricted_area": false})
}

func (h *RestrictedAreaHandler) UpdateArea(c echo.Context) error {
	areaID := c.Param("areaID") // Assuming the area ID is a path parameter

	// Bind request body to a RestrictedArea struct
	var updateArea models.RestrictedArea
	if err := c.Bind(&updateArea); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	// Validate the update data (optional)
	// ... validation logic ...

	// Find the area by ID using the repository
	existingArea, err := h.Repo.FindByID(context.Background(), StrToUint(areaID)) // Replace StrToUint with actual conversion logic
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Area not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Update the existing area with values from updateArea (consider selective updates)
	existingArea.Proponent = updateArea.Proponent // Update specific fields
	existingArea.Branch = updateArea.Branch       // ... and so on

	// Update the area using the repository
	if err := h.Repo.Update(context.Background(), existingArea); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return a successful response
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Area updated successfully"})
}
