package controllers

import (
	"net/http"
	"os"

	"github.com/kapralovs/thinker/internal/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase auth.Usecase
}

type signInRequestBody struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type signInResponse struct {
	Token string `json:"token,omitempty"`
}

func NewAuthHandler(uc auth.Usecase) *AuthHandler {
	return &AuthHandler{usecase: uc}
}

func (h *AuthHandler) SignIn(c echo.Context) error {
	a := &AuthInfo{}
	err := c.Bind(a)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}
	t, err := generateToken(a, os.Getenv("SIGN_STRING"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}

	return c.JSON(http.StatusOK, &signInResponse{token: t})
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	tokenInfo := &TokenInfo{}
	err := c.Bind(tokenInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	return c.JSON(http.StatusCreated, http.StatusText(http.StatusCreated))
}
