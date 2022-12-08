package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kapralovs/thinker/internal/models"
	"github.com/kapralovs/thinker/internal/note"
	"github.com/kapralovs/thinker/internal/utils"

	"github.com/labstack/echo/v4"
)

// type NewNote struct {
// }

type notesHandler struct {
	usecase note.UseCase
}

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
		errMsg := fmt.Sprintf("%s: failed to create a new note: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, utils.ResponseStatusCreated)
}

func (h *notesHandler) EditNote(c echo.Context) error {
	n := new(models.Note)
	err := c.Bind(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}

	err = h.usecase.EditNote(n)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to edit note: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusCreated, utils.ResponseStatusEdited)
}

func (h *notesHandler) DeleteNote(c echo.Context) error {
	strID := c.Param(":id")
	if strID == "" {
		errMsg := fmt.Sprintf("%s: empty path param: %s", utils.ResponseStatusError, "id")
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	id, err := strconv.Atoi(strID)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to parse %s path param: %s", utils.ResponseStatusError, "id", err.Error())
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	err = h.usecase.DeleteNote(int64(id))
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to delete a note: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	return c.JSON(http.StatusOK, utils.ResponseStatusDeleted)
}

func (h *notesHandler) GetNote(c echo.Context) error {
	strID := c.Param(":id")
	if strID == "" {
		errMsg := fmt.Sprintf("%s: empty path param: %s", utils.ResponseStatusError, "id")
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	id, err := strconv.Atoi(strID)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to parse %s path param: %s", utils.ResponseStatusError, "id", err.Error())
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	note, err := h.usecase.GetNote(int64(id))
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to get a note: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	serialized, err := json.Marshal(note)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to marshal response: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, string(serialized))
}

func (h *notesHandler) GetNotesList(c echo.Context) error {
	notes, err := h.usecase.GetNotesList()
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to get a notes list: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	serialized, err := json.Marshal(notes)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to marshal response: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, string(serialized))
}
