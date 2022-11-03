package usecase

import (
	"github.com/kapralovs/thinker/internal/models"
	"github.com/kapralovs/thinker/internal/user"
)

type userUseCase struct {
	repo user.Repository
}

func NewUserUseCase(r user.Repository) *userUseCase {
	return &userUseCase{
		repo: r,
	}
}

func (uc *userUseCase) AddUser(u *models.User) error {
	return nil
}

func (uc *userUseCase) EditUser(id int64) error {
	return nil
}

func (uc *userUseCase) DeleteUser(id int64) error {
	return nil
}

func (uc *userUseCase) GetUser(id int64) (*models.User, error) {
	users, err := uc.repo.GetUser()
	return users, nil
}

func (uc *userUseCase) GetUsersList(id int64) (*models.User, error) {
	users, err := uc.repo.GetUser()
	return users, nil
}
