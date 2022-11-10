package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kapralovs/thinker/internal/models"
	"github.com/kapralovs/thinker/internal/user"
	"github.com/kapralovs/thinker/internal/utils"
	"github.com/labstack/echo/v4"
)

// type NewUser struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

type usersHandler struct {
	usecase user.UseCase
}

func NewUserHandler(uc user.UseCase) *usersHandler {
	return &usersHandler{
		usecase: uc,
	}
}

func (h *usersHandler) CreateUser(c echo.Context) error {
	u := new(models.User)
	err := c.Bind(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}

	err = h.usecase.CreateUser(u)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to create a new note: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, utils.ResponseStatusCreated)
}

func (h *usersHandler) EditUser(c echo.Context) error {
	u := new(models.User)
	err := c.Bind(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}

	err = h.usecase.EditUser(u)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to edit user: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusCreated, utils.ResponseStatusEdited)
}

func (h *usersHandler) DeleteUser(c echo.Context) error {
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

	if err := h.usecase.DeleteUser(int64(id)); err != nil {
		errMsg := fmt.Sprintf("%s: failed to delete a user: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	return c.JSON(http.StatusOK, "success")
}

func (h *usersHandler) GetUser(c echo.Context) error {
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

	user, err := h.usecase.GetUser(int64(id))
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to get a user: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	serialized, err := json.Marshal(user)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to marshal response: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, serialized)
}

func (h *usersHandler) GetUsersList(c echo.Context) error {
	users, err := h.usecase.GetUsersList()
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to get a users list: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	serialized, err := json.Marshal(users)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to marshal response: %s", utils.ResponseStatusError, err.Error())
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, string(serialized))
}
