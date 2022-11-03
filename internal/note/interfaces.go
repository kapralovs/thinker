package note

import "github.com/kapralovs/thinker/internal/models"

type UseCase interface {
	CreateNote(n *models.Note) error
	EditNote(id int64) error
	DeleteNote(id int64) error
	GetNote(id int64) (*models.Note, error)
	GetNotesList() ([]*models.Note, error)
}

type Repository interface {
	CreateNote(n *models.Note) error
	EditNote(id int64) error
	DeleteNote(id int64) error
	GetNote(id int64) (*models.Note, error)
	GetNotesList() ([]*models.Note, error)
}
