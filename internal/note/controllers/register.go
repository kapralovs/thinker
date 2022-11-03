package controllers

import (
	"github.com/kapralovs/thinker/internal/note"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoints(router *echo.Echo, uc note.UseCase) {
	h := NewNoteHandler(uc)

	users := router.Group("/users/")
	users.POST("create", h.CreateNote)
	users.PUT("edit/:id", h.EditNote)
	users.DELETE("delete/:id", h.DeleteNote)
	users.GET("get/:id", h.GetNote)
	users.GET("list", h.GetNotesList)
}
