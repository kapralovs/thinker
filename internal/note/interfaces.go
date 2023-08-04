package note

import "github.com/kapralovs/thinker/internal/models"

type (
	UseCase interface {
		CreateNote(n *models.Note, token *models.AuthClaims) error
		EditNote(n *models.Note, token *models.AuthClaims) error
		DeleteNote(id int64, token *models.AuthClaims) error
		GetNote(id int64, token *models.AuthClaims) (*models.Note, error)
		GetNotesList(filters map[string]string, token *models.AuthClaims) ([]*models.Note, error)
	}

	Repository interface {
		CreateNote(n *models.Note, token *models.AuthClaims) error
		EditNote(u *models.Note, token *models.AuthClaims) error
		DeleteNote(id int64, token *models.AuthClaims) error
		GetNote(id int64, token *models.AuthClaims) (*models.Note, error)
		GetNotesList(filters map[string]string, token *models.AuthClaims) ([]*models.Note, error)
	}
)
