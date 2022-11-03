package controllers

import (
	"net/http"

	"github.com/kapralovs/thinker/internal/user"
	"github.com/labstack/echo/v4"
)

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type usersHandler struct {
	usecase user.UseCase
}

func NewUserHandler(uc user.UseCase) *usersHandler {
	return &usersHandler{
		usecase: uc,
	}
}

func (h *usersHandler) CreateUser(c echo.Context) error {
	u := new(NewUser)
	err := c.Bind(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, "success")
}

// func (h *usersHandler) EditUser(c echo.Context) error {
// 	u := new(NewUser)
// 	err := c.Bind(u)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "bad request")
// 	}
// 	return c.JSON(http.StatusOK, "success")
// }

// func (h *usersHandler) DeleteUser(c echo.Context) error {
// 	u := new(NewUser)
// 	err := c.Bind(u)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "bad request")
// 	}
// 	return c.JSON(http.StatusOK, "success")
// }

// func (h *usersHandler) GetUser(c echo.Context) error {
// 	u := new(NewUser)
// 	err := c.Bind(u)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "bad request")
// 	}
// 	return c.JSON(http.StatusOK, "success")
// }

// func (h *usersHandler) GetUsersList(c echo.Context) error {
// 	u := new(NewUser)
// 	err := c.Bind(u)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "bad request")
// 	}
// 	return c.JSON(http.StatusOK, "success")
// }
