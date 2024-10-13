package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kapralovs/thinker/internal/models"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser() error {

	return nil
}

func (r *UserRepo) GetUser() (*models.User, error) {
	return &models.User{}, nil
}
