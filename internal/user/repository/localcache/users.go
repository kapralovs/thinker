package repository

import (
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
