package models

import (
	"database/sql/driver"
	"fmt"
)

// PostgisGeometry : Geometry type
type PostgisGeometry struct {
	Geometry string
}

// Scan : Implement the `Scanner` interface to read from the database
func (g *PostgisGeometry) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		g.Geometry = v
		return nil
	default:
		return fmt.Errorf("failed to scan PostgisGeometry, expected string, got %T", v)
	}
}

// Value : Implement the `Valuer` interface to write to the database
func (g *PostgisGeometry) Value() (driver.Value, error) {
	return g.Geometry, nil
}

// FlightRestriction : FlightRestriction type
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
