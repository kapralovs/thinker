package controllers

import (
	"github.com/kapralovs/thinker/internal/note"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoints(note *echo.Group, uc note.UseCase) {
	h := NewNoteHandler(uc)
	note.POST("create", h.CreateNote)
	note.PUT("edit/:id", h.EditNote)
	note.DELETE("delete/:id", h.DeleteNote)
	note.GET("get/:id", h.GetNote)
	note.GET("list", h.GetNotesList)
}
