package user

import "github.com/kapralovs/thinker/internal/models"

type UseCase interface {
	CreateUser(u *models.User) error
	EditUser(u *models.User) error
	DeleteUser(id int64) error
	GetUser(id int64) (*models.User, error)
	GetUsersList() ([]*models.User, error)
}

type Repository interface {
	CreateUser(u *models.User) error
	EditUser(u *models.User) error
	DeleteUser(id int64) error
	GetUser(id int64) (*models.User, error)
	GetUsersList() ([]*models.User, error)
}
