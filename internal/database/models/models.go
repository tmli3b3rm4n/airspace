package models

import (
	"database/sql/driver"
	"fmt"
)

type PostgisGeometry struct {
	Geometry string // Store the geometry as WKT (Well-Known Text) or GeoJSON
}

// Implement the `Scanner` interface to read from the database
func (g *PostgisGeometry) Scan(value interface{}) error {
	// Convert the geometry field into a string (WKT or GeoJSON format)
	switch v := value.(type) {
	case string:
		g.Geometry = v
		return nil
	default:
		return fmt.Errorf("failed to scan PostgisGeometry, expected string, got %T", v)
	}
}

// Implement the `Valuer` interface to write to the database
func (g PostgisGeometry) Value() (driver.Value, error) {
	// Return the geometry as a string to be written to the database
	return g.Geometry, nil
}

// Example struct with the custom type
type FlightRestriction struct {
	ID          uint            `gorm:"primaryKey"`
	Proponent   string          `gorm:"column:proponent"`
	Branch      string          `gorm:"column:branch"`
	Base        string          `gorm:"column:base"`
	Facility    string          `gorm:"column:facility"`
	Airspace    string          `gorm:"column:airspace"`
	Reason      string          `gorm:"column:reason"`
	State       string          `gorm:"column:state"`
	FaaID       string          `gorm:"column:faa_id"`
	POC         string          `gorm:"column:poc"`
	Floor       float64         `gorm:"column:floor"`
	Ceiling     float64         `gorm:"column:ceiling"`
	County      string          `gorm:"column:county"`
	ShapeArea   float64         `gorm:"column:shape__area"`
	ShapeLength float64         `gorm:"column:shape__length"`
	Geom        PostgisGeometry `gorm:"type:geometry;column:geom"` // Custom PostGIS geometry type
}
