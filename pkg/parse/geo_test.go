package parse

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestParseLatLon(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]string
		expectedLat   float64
		expectedLon   float64
		expectedError error
	}{
		{
			name:          "Valid parameters",
			params:        map[string]string{"lat": "34.0522", "lon": "-118.2437"},
			expectedLat:   34.0522,
			expectedLon:   -118.2437,
			expectedError: nil,
		},
		{
			name:          "Invalid latitude",
			params:        map[string]string{"lat": "invalid", "lon": "-118.2437"},
			expectedLat:   0,
			expectedLon:   0,
			expectedError: errors.New("flight-restrictions: strconv.ParseFloat: parsing \"invalid\": invalid syntax"),
		},
		{
			name:          "Invalid longitude",
			params:        map[string]string{"lat": "34.0522", "lon": "invalid"},
			expectedLat:   0,
			expectedLon:   0,
			expectedError: errors.New("flight-restrictions: strconv.ParseFloat: parsing \"invalid\": invalid syntax"),
		},
		{
			name:          "Missing parameters",
			params:        map[string]string{},
			expectedLat:   0,
			expectedLon:   0,
			expectedError: errors.New("flight-restrictions: strconv.ParseFloat: parsing \"\": invalid syntax"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := &http.Request{}
			c := e.NewContext(req, nil)

			paramNames := make([]string, 0, len(tt.params))
			paramValues := make([]string, 0, len(tt.params))
			for k, v := range tt.params {
				paramNames = append(paramNames, k)
				paramValues = append(paramValues, v)
			}
			c.SetParamNames(paramNames...)
			c.SetParamValues(paramValues...)

			lat, lon, err := ParseLatLon(c)

			assert.Equal(t, tt.expectedLat, lat)
			assert.Equal(t, tt.expectedLon, lon)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
