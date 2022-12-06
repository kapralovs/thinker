package user

import "github.com/kapralovs/thinker/internal/models"

type Repository interface {
	GetUser(id int64) (*models.User, error)
	GetUsersList() ([]*models.User, error)
	EditUser(u *models.User) error
	DeleteUser(id int64) error
}

type UseCase interface {
	GetUser(id int64) (*models.User, error)
	GetUsersList() ([]*models.User, error)
	EditUser(u *models.User) error
	DeleteUser(id int64) error
}
