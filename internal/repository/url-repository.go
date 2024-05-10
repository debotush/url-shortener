package repository

import (
	"gorm.io/gorm"
	"url-shortener-service/internal/database"
	gormCommonRepository "url-shortener-service/pkg/gorm-common-repository"
)

// UrlsRepository represents a repository for working with URL storage entities.
type UrlsRepository struct {
	DB *gorm.DB
	gormCommonRepository.CommonRepositoryInterface[database.UrlStorage]
}

// NewUrlsRepository creates a new instance of UrlsRepository.
func NewUrlsRepository(DB *gorm.DB) *UrlsRepository {
	return &UrlsRepository{
		DB:                        DB,
		CommonRepositoryInterface: gormCommonRepository.NewCommonRepository[database.UrlStorage]("url_storages", DB),
	}
}
