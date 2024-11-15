package database

import (
	"log"
)

func AutoMigrateAllModels(db Database) error {
	if err := db.AutoMigrate(
		&RestrictedArea{},
	); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
