package util

// IsValidLatitude checks if a float64 is a valid latitude.
func IsValidLatitude(lat float64) bool {
	return lat >= -90.0 && lat <= 90.0
}

// IsValidLongitude checks if a float64 is a valid longitude.
func IsValidLongitude(lon float64) bool {
	return lon >= -180.0 && lon <= 180.0
}
