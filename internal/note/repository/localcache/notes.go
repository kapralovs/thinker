package localcache

import (
	"errors"
	"sync"

	"github.com/kapralovs/thinker/internal/models"
)

type LocalRepo struct {
	notes map[int64]*models.Note
	mu    *sync.Mutex
}

func NewLocalRepo() *LocalRepo {
	return &LocalRepo{
		notes: make(map[int64]*models.Note),
		mu:    new(sync.Mutex),
	}
}

func (r *LocalRepo) CreateNote(n *models.Note) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.notes[n.ID]
	if ok {
		return errors.New("note with such id is already exists")
	}

	r.notes[n.ID] = n

	return nil
}

func (r *LocalRepo) EditNote(n *models.Note) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.notes[n.ID]
	if !ok {
		return errors.New("note with such id is not exists")
	}

	r.notes[n.ID] = n

	return nil
}

func (r *LocalRepo) DeleteNote(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.notes, id)

	return nil
}

func (r *LocalRepo) GetNote(id int64) (*models.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if note, ok := r.notes[id]; ok {
		return note, nil
	}

	return nil, errors.New("заметка с данным ID не найдена")
}

func (r *LocalRepo) GetNotesList() ([]*models.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	notesList := make([]*models.Note, 0, len(r.notes))
	for _, note := range r.notes {
		notesList = append(notesList, note)
	}

	return notesList, nil
}
