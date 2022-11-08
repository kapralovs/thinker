package repository

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

	_, ok := r.users[u.ID]
	if ok {
		return errors.New("note with such id is already exists")
	}

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

func (r *LocalRepo) DeleteUser(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.users, id)

	return nil
}

func (r *LocalRepo) GetUser(id int64) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user, ok := r.users[id]; ok {
		return user, nil
	}

	return nil, errors.New("заметка с данным ID не найдена")
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
