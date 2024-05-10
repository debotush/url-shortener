package database

import "gorm.io/gorm"

// UrlStorage represents the database model for storing URLs.
type UrlStorage struct {
	gorm.Model          // GORM model for managing ID, CreatedAt, UpdatedAt, and DeletedAt fields
	OriginalUrl  string // Original URL
	ShortenerUrl string // Shortened URL
}
