package models

import (
	"gorm.io/gorm"
)

// InitializeDB initializes the database and migrates all models.
func InitializeDB(db *gorm.DB) {
	// Migrate your models here
	db.AutoMigrate(
		&Users{}, 
		// &Product{}
	)

	// Add associations if needed
	SetupAssociations(db)
}

// setupAssociations sets up associations between models.
func SetupAssociations(db *gorm.DB) {
	// Add associations between models here
	// db.Model(&Store{}).Association("Products")
	// db.Model(&Product{}).BelongsTo(&Store{})
}