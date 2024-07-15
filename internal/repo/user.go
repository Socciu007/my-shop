package repo

import "my_shop/internal/models"

// UserRepoer defines methods to interact with user data.
type UserRepoer interface {
	// Create inserts a new user record into the repository.
	Create(u *models.Users) (*models.Users, error)

	// Get retrieves a user record from the repository by ID.
	Get(id string) (*models.Users, error)

	// Update an existing user record in the repository.
	Update(id string, u *models.Users) (*models.Users, error)

	// Delete removes a user record from the repository by its ID.
	Delete(id string) error

	// GetAll retrieves all user records from the repository.
	GetAll() []*models.Users
}