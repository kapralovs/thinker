package auth

import "github.com/kapralovs/thinker/internal/models"

type (
	Repository interface {
		GetUser(username, password string) (*models.User, error)
		CreateUser(user *models.User) error
	}

	UseCase interface {
		SignIn(username, password string) (string, error)
		SignUp(username, password string) (string, error)
		ParseToken(token string) (*models.AuthClaims, error)
	}
)
