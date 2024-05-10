package gorm_common_repository

// CommonRepositoryInterface defines the common methods that a repository should implement.
type CommonRepositoryInterface[Model any] interface {

	// Store stores a data object in the database.
	Store(dataObject Model) (Model, error)

	// FindBy finds a data object in the database by a specified field and value.
	FindBy(field string, value any) (Model, error)
}