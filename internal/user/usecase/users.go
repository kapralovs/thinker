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

func (uc *userUseCase) CreateUser(u *models.User) error {
	err := uc.repo.CreateUser(u)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) EditUser(u *models.User) error {
	err := uc.repo.EditUser(u)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) DeleteUser(id int64) error {
	err := uc.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) GetUser(id int64) (*models.User, error) {
	note, err := uc.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (uc *userUseCase) GetUsersList() ([]*models.User, error) {
	users, err := uc.repo.GetUsersList()
	if err != nil {
		return nil, err
	}

	return users, nil
}
