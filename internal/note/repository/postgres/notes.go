package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kapralovs/thinker/internal/models"
)

type NoteRepo struct {
	db *pgxpool.Pool
}

func NewNoteRepo(db *pgxpool.Pool) *NoteRepo {
	return &NoteRepo{
		db: db,
	}
}

func (nr *NoteRepo) CreateNote() error {

	return nil
}

func (nr *NoteRepo) UpdateNote() error {

	return nil

}

func (nr *NoteRepo) DeleteNote() error {

	return nil

}

func (nr *NoteRepo) GetNote() (*models.Note, error) {
	return &models.Note{}, nil
}
