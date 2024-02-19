package user

import (
	"errors"
	"sync"

	"github.com/kapralovs/thinker/internal/models"
)

type UserLocalRepo struct {
	mu    *sync.Mutex
	users map[int64]*models.User
}

func NewUserLocalRepo() *UserLocalRepo {
	return &UserLocalRepo{
		mu:    new(sync.Mutex),
		users: make(map[int64]*models.User),
	}
}

func (r *UserLocalRepo) GetUser(id int64) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if u, ok := r.users[id]; ok {
		return u, nil
	}

	return nil, errors.New("user is not exist")
}

func (r *UserLocalRepo) GetUsersList() ([]*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	usersList := []*models.User{}
	for _, u := range r.users {
		usersList = append(usersList, u)
	}

	return usersList, nil
}

func (r *UserLocalRepo) EditUser(u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[u.ID]; ok {
		r.users[u.ID] = u
	}

	return errors.New("user is not exist")
}

func (r *UserLocalRepo) DeleteUser(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.users, id)

	return nil
}
