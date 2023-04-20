package note

import "github.com/kapralovs/thinker/internal/models"

type (
	UseCase interface {
		CreateNote(n *models.Note) error
		EditNote(n *models.Note) error
		DeleteNote(id int64) error
		GetNote(id int64) (*models.Note, error)
		GetNotesList() ([]*models.Note, error)
	}

	Repository interface {
		CreateNote(n *models.Note) error
		EditNote(u *models.Note) error
		DeleteNote(id int64) error
		GetNote(id int64) (*models.Note, error)
		GetNotesList() ([]*models.Note, error)
	}
)
