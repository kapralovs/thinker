package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kapralovs/thinker/internal/models"
	"github.com/kapralovs/thinker/internal/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase user.UseCase
}

func NewUserHandler(uc user.UseCase) *UserHandler {
	return &UserHandler{
		usecase: uc,
	}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("can't get user: %s", err.Error()))
	}

	u, err := h.usecase.GetUser(int64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't get user: %s", err.Error()))
	}

	return c.JSON(http.StatusOK, u)
}

func (h *UserHandler) GetUsersList(c echo.Context) error {
	users, err := h.usecase.GetUsersList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't get users list: %s", err.Error()))
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) EditUser(c echo.Context) (err error) {
	u := new(models.User)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("can't get user: %s", err.Error()))
	}

	if err = h.usecase.EditUser(u); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't edit user: %s", err.Error()))
	}

	return c.JSON(http.StatusOK, "edited")
}
