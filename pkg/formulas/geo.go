package formulas

import "math"

// Haversine calculates the distance between two points on a sphere (Earth) in meters.
//
// The function takes four float64 arguments:
//   - lat1: Latitude of the first point in decimal degrees.
//   - lon1: Longitude of the first point in decimal degrees.
//   - lat2: Latitude of the second point in decimal degrees.
//   - lon2: Longitude of the second point in decimal degrees.
//
// The function returns a float64 value representing the distance in meters.
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad

	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
