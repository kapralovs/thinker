package controllers

import (
	"github.com/kapralovs/thinker/internal/note"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoints(note *echo.Group, uc note.UseCase) {
	h := NewNoteHandler(uc)

	note.POST("create", h.createNote)
	note.PUT("edit/:id", h.editNote)
	note.DELETE("delete/:id", h.deleteNote)
	note.GET(":id", h.getNote)
	note.GET("list", h.getNotesList)
}
