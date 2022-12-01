package localcache

import (
	"errors"
	"sync"

	"github.com/kapralovs/thinker/internal/models"
)

type LocalRepo struct {
	mu    *sync.Mutex
	users map[string]*models.User
}

func NewLocalRepo() *LocalRepo {
	return &LocalRepo{
		mu:    new(sync.Mutex),
		users: make(map[string]*models.User),
	}
}

func (r *LocalRepo) CreateUser(u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.users[u.Username]
	if ok {
		return errors.New("note with such id is already exists")
	}

	r.users[u.Username] = &models.User{Username: u.Username, Password: u.Username}

	return nil
}

func (r *LocalRepo) EditUser(u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.users[u.Username]
	if !ok {
		return errors.New("note with such id is not exists")
	}

	r.users[u.Username] = u

	return nil
}

// func (r *LocalRepo) DeleteUser(id int64) error {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	delete(r.users, u.Username)

// 	return nil
// }

func (r *LocalRepo) GetUser(username, password string) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user, ok := r.users[username]; ok {
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
