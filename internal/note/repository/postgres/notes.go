package postgres

import (
	"context"

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

func (nr *NoteRepo) CreateNote(n *models.Note, token *models.AuthClaims) error {
	if err := nr.db.QueryRow(context.Background(), "INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id",
		n.Title,
		n.Text,
	).Scan(&n.ID); err != nil {
		return err
	}

	return nil
}

func (nr *NoteRepo) EditNote(n *models.Note, token *models.AuthClaims) error {
	if err := nr.db.QueryRow(context.Background(), "UPDATE notes SET title = $1, content = $2 WHERE id = $3",
		n.Title,
		n.Text,
	).Scan(&n.ID); err != nil {
		return err
	}

	return nil
}

func (nr *NoteRepo) DeleteNote(id int64, token *models.AuthClaims) error {

	return nil

}

func (nr *NoteRepo) GetNote(id int64, token *models.AuthClaims) (*models.Note, error) {
	return &models.Note{}, nil
}

func (nr *NoteRepo) GetNotesList(filters map[string]string, token *models.AuthClaims) ([]*models.Note, error) {
	return []*models.Note{}, nil
}
