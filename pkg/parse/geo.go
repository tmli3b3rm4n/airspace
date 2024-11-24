package parse

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParseLatLon(c echo.Context) (float64, float64, error) {
	lat, err := strconv.ParseFloat(c.Param("lat"), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("flight-restrictions: %v", err)
	}

	lon, err := strconv.ParseFloat(c.Param("lon"), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("flight-restrictions: %v", err)
	}
	return lat, lon, nil
}
