package gorm_common_repository

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CommonRepository is a generic repository implementation for common database operations.
type CommonRepository[Model any] struct {
	GenericStruct Model    // The model structure for the database entity.
	TableName     string   // The name of the database table.
	DB            *gorm.DB // The GORM database instance.
}

// NewCommonRepository creates a new instance of CommonRepository.
func NewCommonRepository[Model any](tableName string, DB *gorm.DB) CommonRepositoryInterface[Model] {
	return &CommonRepository[Model]{
		TableName: tableName,
		DB:        DB,
	}
}

// Store stores a data object in the database.
func (repo *CommonRepository[Model]) Store(dataObject Model) (Model, error) {

	// Execute a database query to store the data object in the database
	err := repo.DB.Clauses(clause.Returning{}).Table(repo.TableName).Save(&dataObject).Error

	if err != nil {
		// If an error occurs during the database operation, return the data object and the error
		return dataObject, err
	}

	// Return the stored data object and nil error
	return dataObject, nil
}

// FindBy finds a data object in the database by a specified field and value.
func (repo *CommonRepository[Model]) FindBy(field string, value any) (Model, error) {
	var dataObject Model

	// Construct a WHERE clause for the specified field and value
	whereQ := fmt.Sprintf("%s = ?", field)

	// Execute the database query to find the data object
	err := repo.DB.Table(repo.TableName).Where(whereQ, value).First(&dataObject).Error

	if err != nil {
		// If an error occurs, return an empty data object and the error
		return dataObject, err
	}

	// Return the found data object and nil error
	return dataObject, nil
}
