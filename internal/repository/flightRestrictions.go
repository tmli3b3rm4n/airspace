package repository

import (
	"fmt"
	"github.com/tmli3b3rm4n/airspace/internal/database"
	"github.com/tmli3b3rm4n/airspace/internal/database/models"
)

// IFlightRestrictions defines the interface for interacting with flight restriction data
type IFlightRestrictions interface {
	// RestrictedAirspace checks if the given lat/lon is within restricted airspace
	RestrictedAirspace(lat, long float64) (bool, error)
}

// NewFlightRestrictionsRepo creates a new instance of FlightRestrictionsRepo
func NewFlightRestrictionsRepo(db database.Database) IFlightRestrictions {
	return &FlightRestrictionsRepo{
		db: db,
	}
}

// FlightRestrictionsRepo is the struct for handling flight restriction queries
type FlightRestrictionsRepo struct {
	db database.Database
}

// RestrictedAirspace checks if the provided coordinates are inside restricted airspace
func (f *FlightRestrictionsRepo) RestrictedAirspace(lat, lon float64) (bool, error) {
	var count int64

	// Perform the query to check if the point is within restricxxted airspace
	if err := f.db.Model(&models.FlightRestriction{}).Where(
		"ST_Intersects(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326))", lon, lat).
		Count(&count).Error; err != nil {
		return false, fmt.Errorf("failed to query database for point (%f, %f): %v", lat, lon, err)
	}

	// If count > 0, the point is inside restricted airspace
	if count > 0 {
		return true, nil
	}
	return false, nil
}
