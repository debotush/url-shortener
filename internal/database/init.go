package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	db  *gorm.DB
	err error
)

// InitializeGormDatabase initializes a GORM database connection.
// It configures GORM, opens a connection to the database, and performs database setup.
func InitializeGormDatabase() *gorm.DB {
	// Database connection URL
	dbURL := "postgres://postgres:r00t@localhost:54321/url_shortener?sslmode=disable"

	// GORM configuration options
	gormConfig := configureGorm()

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dbURL), gormConfig)
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}

	// Set maximum idle and open connections
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)

	// Perform database setup
	setupDatabase(db)

	return db
}

// configureGorm configures GORM with custom options.
func configureGorm() *gorm.Config {
	return &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // IO writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Enable color
			},
		),
	}
}

// setupDatabase performs necessary database setup, including auto-migration and connection pool settings.
func setupDatabase(db *gorm.DB) {
	// Auto-migrate tables if necessary
	db.AutoMigrate(&UrlStorage{})
}

// GetDbInstance returns the instance of the GORM database.
func GetDbInstance() *gorm.DB {
	return db
}
