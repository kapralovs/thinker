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

type AuthResponse struct {
	Token   string `json:"access_token,omitempty"`
	Success bool   `json:"success,omitempty"`
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

	ar := &AuthResponse{Token: token, Success: true}

	return c.JSON(http.StatusOK, ar)
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	s := new(signUpRequestBody)
	err := c.Bind(s)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	token, err := h.usecase.SignUp(s.Username, s.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "sign up error")
	}

	ar := &AuthResponse{Token: token, Success: true}

	return c.JSON(http.StatusCreated, ar)
}
