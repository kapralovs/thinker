package usecase

import (
	"github.com/kapralovs/thinker/internal/models"
	"github.com/kapralovs/thinker/internal/user"
)

type UserUseCase struct {
	repo user.Repository
}

func NewUserUseCase(r user.Repository) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) GetUser(id int64) (*models.User, error) {
	u, err := uc.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uc *UserUseCase) GetUsersList() ([]*models.User, error) {
	u, err := uc.repo.GetUsersList()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uc *UserUseCase) EditUser(u *models.User) error {
	err := uc.repo.EditUser(u)
	if err != nil {
		return err
	}

	return nil
}
