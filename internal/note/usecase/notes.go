package usecase

import (
	"github.com/kapralovs/thinker/internal/models"
	"github.com/kapralovs/thinker/internal/note"
)

type notesUseCase struct {
	repo note.Repository
}

func NewNoteUseCase(r note.Repository) *notesUseCase {
	return &notesUseCase{
		repo: r,
	}
}

func (uc *notesUseCase) CreateNote(n *models.Note) error {
	err := uc.repo.CreateNote(n)
	if err != nil {
		return err
	}

	return nil
}

func (uc *notesUseCase) EditNote(n *models.Note) error {
	err := uc.repo.EditNote(n)
	if err != nil {
		return err
	}

	return nil
}

func (uc *notesUseCase) DeleteNote(id int64) error {
	err := uc.repo.DeleteNote(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *notesUseCase) GetNote(id int64) (*models.Note, error) {
	note, err := uc.repo.GetNote(id)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (uc *notesUseCase) GetNotesList() ([]*models.Note, error) {
	notes, err := uc.repo.GetNotesList()
	if err != nil {
		return nil, err
	}

	return notes, nil
}
