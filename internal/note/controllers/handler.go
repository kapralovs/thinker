package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kapralovs/thinker/internal/models"
	"github.com/kapralovs/thinker/internal/note"
	"github.com/labstack/echo/v4"
)

// type NewNote struct {
// }

type notesHandler struct {
	usecase note.UseCase
}

const (
	ResponseStatusCreated = "created"
	ResponseStatusEdited  = "edited"
	ResponseStatusDeleted = "deleted"
)

func NewNoteHandler(uc note.UseCase) *notesHandler {
	return &notesHandler{
		usecase: uc,
	}
}

func (h *notesHandler) CreateNote(c echo.Context) error {
	n := new(models.Note)
	err := c.Bind(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}

	err = h.usecase.CreateNote(n)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to create a new note")
	}

	return c.JSON(http.StatusOK, ResponseStatusCreated)
}

func (h *notesHandler) EditNote(c echo.Context) error {
	n := new(models.Note)
	err := c.Bind(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}

	err = h.usecase.EditNote(n.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to create a new note")
	}

	return c.JSON(http.StatusCreated, ResponseStatusEdited)
}

func (h *notesHandler) DeleteNote(c echo.Context) error {
	strID := c.Param(":id")
	if strID == "" {
		errMsg := fmt.Sprintf("empty path param: %s", "id")
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	id, err := strconv.Atoi(strID)
	if err != nil {
		errMsg := fmt.Sprintf("failed to parse %s path param: %s", "id", err.Error())
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	err = h.usecase.DeleteNote(int64(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, ResponseStatusDeleted)
}
