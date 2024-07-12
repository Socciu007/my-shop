package models

import (
	"gorm.io/gorm"
)

// InitializeDB initializes the database and migrates all models.
func InitializeDB(db *gorm.DB) error {
	// Migrate each model
	if err := db.AutoMigrate(
		&Users{},
		// &Product{},
		// &Order{},
	); err != nil {
		return err
	}

	// Setup model associations
	if err := setupAssociations(db); err != nil {
		return err
	}

	return nil
}

// setupAssociations sets up associations between models.
func setupAssociations(db *gorm.DB) error {
	// Add associations between models here
	return nil
}
