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

func (uc *notesUseCase) CreateNote(n *models.Note, token *models.AuthClaims) (err error) {
	if err := uc.repo.CreateNote(n, token); err != nil {
		return err
	}

	return nil
}

func (uc *notesUseCase) EditNote(n *models.Note, token *models.AuthClaims) (err error) {
	if err := uc.repo.EditNote(n, token); err != nil {
		return err
	}

	return nil
}

func (uc *notesUseCase) DeleteNote(id int64, token *models.AuthClaims) (err error) {
	if err := uc.repo.DeleteNote(id, token); err != nil {
		return err
	}

	return nil
}

func (uc *notesUseCase) GetNote(id int64, token *models.AuthClaims) (*models.Note, error) {
	note, err := uc.repo.GetNote(id, token)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (uc *notesUseCase) GetNotesList(filters map[string]string, token *models.AuthClaims) ([]*models.Note, error) {
	notes, err := uc.repo.GetNotesList(filters, token)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
