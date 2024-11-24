package cords

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tmli3b3rm4n/airspace/pkg/util"
	"strconv"
)

func GetLatLon(c echo.Context) (float64, float64, error) {
	lat, err := strconv.ParseFloat(c.Param("lat"), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("flight-restrictions: %v", err)
	}

	lon, err := strconv.ParseFloat(c.Param("lon"), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("flight-restrictions: %v", err)
	}

	if !util.IsValidLatitude(lat) {
		return 0, 0, fmt.Errorf("invalid latitude")
	}
	if !util.IsValidLongitude(lon) {
		return 0, 0, fmt.Errorf("Invalid longitude")
	}
	return lat, lon, nil
}
