package localcache

import (
	"errors"
	"sync"

	"github.com/kapralovs/thinker/internal/models"
)

type LocalRepo struct {
	mu    *sync.Mutex
	users map[int64]*models.User
}

func NewLocalRepo() *LocalRepo {
	return &LocalRepo{
		mu:    new(sync.Mutex),
		users: make(map[int64]*models.User),
	}
}

func (r *LocalRepo) CreateUser(u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if u.Username == user.Username {
			return errors.New("user with such username is already exists")
		}
	}

	u.ID = int64(len(r.users) + 1)
	r.users[u.ID] = u

	return nil
}

func (r *LocalRepo) EditUser(u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.users[u.ID]
	if !ok {
		return errors.New("note with such id is not exists")
	}

	r.users[u.ID] = u

	return nil
}

func (r *LocalRepo) GetUser(username, password string) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Username == username {
			if user.Password == password {
				return user, nil
			}
		}
	}

	return nil, errors.New("пользователь с данным ID не найден")
}

func (r *LocalRepo) GetUsersList() ([]*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	userList := make([]*models.User, 0, len(r.users))
	for _, user := range r.users {
		userList = append(userList, user)
	}

	return userList, nil
}
