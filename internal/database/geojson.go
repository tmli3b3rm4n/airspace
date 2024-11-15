package database

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

// ParseGeoJSON reads a GeoJSON file and inserts it into the database
func ParseGeoJSON(db *gorm.DB, filePath string) error {
	// Open the GeoJSON file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read file content
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// Unmarshal JSON data into GeoJSON struct
	var geoData GeoJSON
	if err := json.Unmarshal(byteValue, &geoData); err != nil {
		return err
	}

	// Save GeoJSON structure to the database
	for _, feature := range geoData.Features {
		// Save Property if it does not exist already
		if err := db.Where(Property{ObjectID: feature.Property.ObjectID}).
			Attrs(feature.Property).
			FirstOrCreate(&feature.Property).Error; err != nil {
			return fmt.Errorf("failed to save property: %w", err)
		}

		// Associate Property with the feature
		feature.GeoJSONID = geoData.ID
		if err := db.Create(&feature).Error; err != nil {
			return fmt.Errorf("failed to save geoJSON feature: %w", err)
		}
	}
	return nil
}
