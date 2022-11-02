package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type handler struct {
	usecase users.Usecase
}

func NewHandler(uc) *handler {
	return &handler{
		usecase: uc,
	}
}

func (h *handler) CreateUser(c echo.Context) error {
	u := new(NewUser)
	err := c.Bind(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}

}
