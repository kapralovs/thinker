package users

import "github.com/kapralovs/thinker/internal/models"

type UseCase interface {
	AddUser(n *models.User) error
	EditUser(id int64) error
	DeleteUser(id int64) error
	ViewUser(id int64) (*models.User, error)
}

type Repository interface {
	CreateUser(n *models.User) error
	EditUser(id int64) error
	DeleteUser(id int64) error
	GetUser(id int64) (*models.User, error)
}
