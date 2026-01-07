package database

import (
	"log"

	"github.com/InvertedBit/hl4b-portal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect establishes a connection to the PostgreSQL database
func Connect(databaseURL string) {
	var err error

	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")
}

// AutoMigrate runs automatic migration for the portal tables
// Note: This does NOT migrate the auth.users table which is managed externally
func AutoMigrate() {
	log.Println("Running auto-migration...")

	// Only migrate our custom tables, not auth.users
	err := DB.AutoMigrate(
		&models.PortalUser{},
		&models.Upload{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
