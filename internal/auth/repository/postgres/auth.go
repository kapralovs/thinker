package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kapralovs/thinker/internal/models"
)

type UserRepo struct {
	db *pgxpool.Pool
}

const (
	dbCreateUser = `
	INSERT INTO users (
		name,
		username,
		password,
		current_token
	) VALUES ($1, $2, $3, $4) 
	RETURNING id;`

	dbUpdateUser = `
	UPDATE users 
	SET
		current_token=$1
	WHERE id=$2
	RETURNING id;`

	dbGetUser = `
	SELECT 
		id,
		name,
		username,
		password,
		current_token
	FROM users 
	WHERE username = $1 
	AND password = $2;
	`
)

func NewAuthRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: pool}
}

func (r *UserRepo) CreateUser(u *models.User) error {
	if err := r.db.QueryRow(context.Background(), dbCreateUser,
		u.Username,
		u.Username,
		u.Password,
		u.CurrentToken,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) UpdateUser(u *models.User) error {
	if err := r.db.QueryRow(context.Background(), dbUpdateUser,
		u.CurrentToken,
		u.ID,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetUser(username, password string) (*models.User, error) {
	u := &models.User{
		Username: username,
		Password: password,
	}

	if err := r.db.QueryRow(context.Background(), dbGetUser,
		u.Username,
		u.Password,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Username,
		&u.Password,
		&u.CurrentToken,
	); err != nil {
		return nil, err
	}
	return u, nil
}
