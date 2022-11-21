package controllers

import (
	"net/http"

	"github.com/kapralovs/thinker/internal/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase auth.Usecase
}

type signInRequestBody struct {
	username string `json:"username,omitempty"`
	password string `json:"password,omitempty"`
}

type signInResponse struct {
	token string `json:"token,omitempty"`
}

func NewAuthHandler(uc auth.Usecase) *AuthHandler {
	return &AuthHandler{usecase: uc}
}

func (h *AuthHandler) SignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, &signInResponse{token: tokenString})
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	tokenInfo := &TokenInfo{}
	err := c.Bind(tokenInfo)
	return c.JSON(http.StatusCreated, http.StatusText(http.StatusCreated))
}
