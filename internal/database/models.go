package database

import (
	"gorm.io/gorm"
)

// Property structure for restricted area properties
type Property struct {
	gorm.Model
	ObjectID      int     `json:"OBJECTID" gorm:"column:OBJECTID;uniqueIndex"` // Unique identifier
	Proponent     string  `json:"Proponent"`
	Branch        string  `json:"Branch"`
	Base          string  `json:"Base"`
	Facility      string  `json:"Facility"`
	Airspace      string  `json:"Airspace"`
	Reason        string  `json:"Reason"`
	State         string  `json:"State"`
	FAA_ID        string  `json:"FAA_ID"`
	POC           string  `json:"POC"`
	Floor         string  `json:"Floor"`
	Ceiling       string  `json:"Ceiling"`
	County        string  `json:"County"`
	Shape__Area   float64 `json:"Shape__Area"`
	Shape__Length float64 `json:"Shape__Length"`
	Restricted    bool
}

// GeoJSON structure for parsing GeoJSON data
type GeoJSON struct {
	gorm.Model
	Type     string           `json:"type"`
	Features []GeoJSONFeature `json:"features" gorm:"foreignKey:GeoJSONID"`
}

// GeoJSONFeature structure representing each feature in GeoJSON
type GeoJSONFeature struct {
	gorm.Model
	GeoJSONID uint     `gorm:"index"` // Foreign key to associate with GeoJSON
	Type      string   `json:"type"`
	Geometry  Geometry `json:"geometry" gorm:"embedded;embeddedPrefix:geometry_"`
	Property  Property `json:"properties" gorm:"foreignKey:ObjectID;references:ObjectID"`
}

// Geometry structure for polygon with 3D coordinates
type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates" gorm:"type:json"` // For 3D polygon coordinates
}
