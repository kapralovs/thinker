package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kapralovs/thinker/internal/models"
	"github.com/kapralovs/thinker/internal/note"
	"github.com/kapralovs/thinker/internal/utils"

	"github.com/labstack/echo/v4"
)

type notesHandler struct {
	usecase note.UseCase
}

func NewNoteHandler(uc note.UseCase) *notesHandler {
	return &notesHandler{usecase: uc}
}

func (h *notesHandler) createNote(c echo.Context) (err error) {
	n := new(models.Note)

	if err = c.Bind(n); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}

	token := c.Request().Context().Value("token").(*models.AuthClaims)

	if err = h.usecase.CreateNote(n, token); err != nil {
		errMsg := fmt.Sprintf("%s: failed to create a new note: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, utils.ResponseStatusCreated)
}

func (h *notesHandler) editNote(c echo.Context) (err error) {
	n := new(models.Note)

	if err = c.Bind(n); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}

	token := c.Request().Context().Value("token").(*models.AuthClaims)

	if err = h.usecase.EditNote(n, token); err != nil {
		errMsg := fmt.Sprintf("%s: failed to edit note: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusCreated, utils.ResponseStatusEdited)
}

func (h *notesHandler) deleteNote(c echo.Context) error {
	idParam := c.Param(":id")

	if idParam == "" {
		errMsg := fmt.Sprintf("%s: empty path param: %s", utils.ResponseStatusError, "id")
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to parse %s path param: %s", utils.ResponseStatusError, "id", err.Error())
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	token := c.Request().Context().Value("token").(*models.AuthClaims)

	if err = h.usecase.DeleteNote(int64(id), token); err != nil {
		errMsg := fmt.Sprintf("%s: failed to delete a note: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	return c.JSON(http.StatusOK, utils.ResponseStatusDeleted)
}

func (h *notesHandler) getNote(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to parse %s path param: %s", utils.ResponseStatusError, "id", err.Error())
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	token := c.Request().Context().Value("token").(*models.AuthClaims)

	note, err := h.usecase.GetNote(int64(id), token)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to get a note: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, note)
}

func (h *notesHandler) getNotesList(c echo.Context) error {
	tagParam := c.QueryParam("tag")
	filters := make(map[string]string, 0)

	if tagParam != "" {
		filters["tag"] = tagParam
	}

	token := c.Request().Context().Value("token").(*models.AuthClaims)

	notes, err := h.usecase.GetNotesList(filters, token)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to get a notes list: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, notes)
}
