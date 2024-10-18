package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kapralovs/thinker/internal/models"
)

type userRepo struct {
	db *pgxpool.Pool
}

const (
	dbGetUser = `
	SELECT
		u.id,
		u.name,
		u.username
	FROM users u
	WHERE u.id = $1;`

	dbGetUsersList = `
	SELECT
		u.id,
		u.name,
		u.username
	FROM users u;`

	dbEditUser = `
	UPDATE users
	SET name = $1
	WHERE u.id = $2;`

	dbDeleteUser = `
	DELETE FROM users
	WHERE id = $1;`
)

func NewUserRepo(conn *pgxpool.Pool) *userRepo {
	return &userRepo{db: conn}
}

func (r *userRepo) GetUser(id int64) (*models.User, error) {
	u := &models.User{}

	if err := r.db.QueryRow(context.Background(), dbGetUser, id).Scan(
		&u.ID,
		&u.Name,
		&u.Username,
	); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepo) GetUsersList() ([]models.User, error) {
	rows, err := r.db.Query(context.Background(), dbGetUsersList)
	if err != nil {
		return nil, err
	}

	var (
		u     models.User
		users []models.User
	)

	for rows.Next() {
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Username,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *userRepo) EditUser(u *models.User) error {
	if err := r.db.QueryRow(context.Background(), dbEditUser,
		u.Name,
		u.ID,
	).Scan(&u.ID); err != nil {
		return err
	}

	return errors.New("user is not exist")
}

func (r *userRepo) DeleteUser(id int64) error {
	if err := r.db.QueryRow(context.Background(), dbDeleteUser,
		id,
	).Scan(&id); err != nil {
		return err
	}

	return nil
}
