package auth

import "github.com/kapralovs/thinker/internal/models"

type Repository interface {
	GetUser(username, password string) (*models.User, error)
	CreateUser(user *models.User) error
}

type Usecase interface {
	SignIn(username, password string) (string, error)
	SignUp(user, password string) error
}
