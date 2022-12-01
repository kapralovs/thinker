package controllers

import (
	"net/http"

	"github.com/kapralovs/thinker/internal/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase auth.UseCase
}

type signInRequestBody struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type signUpRequestBody struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

type signInResponse struct {
	Token string `json:"token,omitempty"`
}

func NewAuthHandler(uc auth.UseCase) *AuthHandler {
	return &AuthHandler{usecase: uc}
}

func (h *AuthHandler) SignIn(c echo.Context) error {
	a := new(AuthInfo)
	err := c.Bind(a)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	token, err := h.usecase.SignIn(a.Login, a.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "can't sign in")
	}

	return c.JSON(http.StatusOK, token)
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	a := &AuthInfo{}
	err := c.Bind(a)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	if err = h.usecase.SignUp(a.Login, a.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, "can't sign in")
	}
	//TODO: закончить авторизацию

	return c.JSON(http.StatusCreated, "created")
}
