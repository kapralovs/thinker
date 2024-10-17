package usecase

import (
	"github.com/kapralovs/thinker/internal/models"
	"github.com/kapralovs/thinker/internal/user"
)

type userUseCase struct {
	repo user.Repository
}

func NewuserUseCase(r user.Repository) *userUseCase {
	return &userUseCase{repo: r}
}

func (uc *userUseCase) GetUser(id int64) (*models.User, error) {
	u, err := uc.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uc *userUseCase) GetUsersList() ([]models.User, error) {
	u, err := uc.repo.GetUsersList()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uc *userUseCase) EditUser(u *models.User) error {
	if err := uc.repo.EditUser(u); err != nil {
		return err
	}

	return nil
}
