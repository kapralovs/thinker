package localcache

import (
	"errors"
	"fmt"
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

	n.ID = int64(len(r.notes) + 1)
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

func (r *LocalRepo) GetNotesList(filters map[string][]string) ([]*models.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	notesList := make([]*models.Note, 0)
	for _, note := range r.notes {
		if len(filters) > 0 {
			if applyFilters(filters, note) {
				notesList = append(notesList, note)
			}

			continue
		}

		notesList = append(notesList, note)
	}

	return notesList, nil
}

func applyFilters(filters map[string][]string, note *models.Note) bool {
	for fName, fValues := range filters {
		switch fName {
		case "tags":
			for _, required := range fValues {
				fmt.Println(required)
				if !checkTag(required, note.Tags) {
					return false
				}
			}
		}
	}

	return true
}

func checkTag(required string, tags []string) bool {
	for _, tag := range tags {
		if tag == required {
			return true
		}
	}

	return false
}
