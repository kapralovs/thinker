package postgres

import (
	"context"

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

func (r *UserRepo) CreateUser(user *models.User) error {
	if err := r.db.QueryRow(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id",
		user.Username,
		user.Password,
	).Scan(&user.ID); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetUser(username, password string) (*models.User, error) {
	user := new(models.User)
	if err := r.db.QueryRow(context.Background(), "SELECT * FROM users WHERE username = $1 AND password = $2",
		user.Username,
		user.Password,
	).Scan(&user.ID); err != nil {
		return nil, err
	}
	return &models.User{}, nil
}
