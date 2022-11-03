package usecase

import (
	"github.com/kapralovs/thinker/internal/models"
	"github.com/kapralovs/thinker/internal/note"
)

type NotesUsecase struct {
	repo note.Repository
}

func NewNotesUsecase(r note.Repository) *NotesUsecase {
	return &NotesUsecase{
		repo: r,
	}
}

func (uc *NotesUsecase) CreateNote(n *models.Note) error {
	err := uc.repo.CreateNote(n)
	if err != nil {
		return err
	}

	return nil
}

func (uc *NotesUsecase) EditNote(id int64) error {
	err := uc.repo.EditNote(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *NotesUsecase) DeleteNote(id int64) error {
	err := uc.repo.DeleteNote(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *NotesUsecase) GetNote(id int64) (*models.Note, error) {
	note, err := uc.repo.GetNote(id)
	if err != nil {
		return nil, err
	}

	return note, nil
}
