package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kapralovs/thinker/internal/models"
)

type NoteRepo struct {
	db *pgxpool.Pool
}

const (
	dbGetNote = `
	SELECT 
		id, 
		created_by,
		title, 
		text
	FROM notes 
	WHERE id = $1`

	dbGetNotes = `
	SELECT
		id,
		created_by,
		title,
		text
	FROM notes`

	dbCreateNote = `
	INSERT INTO notes (
		created_by,
		title,
		text
	) VALUES ($1, $2, $3) 
	RETURNING id`

	dbDeleteNote = `
	DELETE FROM notes 
	WHERE id = $1`
)

func NewNoteRepo(db *pgxpool.Pool) *NoteRepo {
	return &NoteRepo{
		db: db,
	}
}

func (nr *NoteRepo) CreateNote(n *models.Note, token *models.AuthClaims) error {
	if err := nr.db.QueryRow(context.Background(), dbCreateNote,
		token.User.Username,
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
	if err := nr.db.QueryRow(context.Background(), dbDeleteNote, id).Scan(
		&id,
	); err != nil {
		return err
	}

	return nil
}

func (nr *NoteRepo) GetNote(id int64, token *models.AuthClaims) (*models.Note, error) {
	n := &models.Note{ID: id}

	if err := nr.db.QueryRow(context.Background(), dbGetNote, n.ID).Scan(
		&n.ID,
		&n.Title,
		&n.Text,
	); err != nil {
		return nil, err
	}

	return n, nil
}

func (nr *NoteRepo) GetNotesList(filters map[string]string, token *models.AuthClaims) ([]models.Note, error) {
	rows, err := nr.db.Query(context.Background(), dbGetNotes)
	if err != nil {
		return nil, err
	}

	var (
		n     models.Note
		notes []models.Note
	)

	for rows.Next() {
		if err := rows.Scan(
			&n.ID,
			&n.CreatedBy,
			&n.Title,
			&n.Text,
		); err != nil {
			return nil, err
		}

		notes = append(notes, n)
	}

	return notes, nil
}
