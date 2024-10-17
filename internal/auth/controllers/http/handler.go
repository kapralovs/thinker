package controllers

import (
	"net/http"

	"github.com/kapralovs/thinker/internal/auth"
	"github.com/labstack/echo/v4"
)

type (
	authHandler struct {
		usecase auth.UseCase
	}

	signUpRequestBody struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Name     string `json:"name,omitempty"`
	}

	AuthResponse struct {
		Token   string `json:"access_token,omitempty"`
		Success bool   `json:"success,omitempty"`
	}
)

func NewAuthHandler(uc auth.UseCase) *authHandler {
	return &authHandler{usecase: uc}
}

func (h *authHandler) signIn(c echo.Context) error {
	authInfo := new(AuthInfo)

	if err := c.Bind(authInfo); err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	token, err := h.usecase.SignIn(authInfo.Login, authInfo.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	authResp := &AuthResponse{Token: token, Success: true}

	return c.JSON(http.StatusOK, authResp)
}

func (h *authHandler) signUp(c echo.Context) error {
	signUpReq := new(signUpRequestBody)

	if err := c.Bind(signUpReq); err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}

	token, err := h.usecase.SignUp(signUpReq.Username, signUpReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	authResp := &AuthResponse{Token: token, Success: true}

	return c.JSON(http.StatusCreated, authResp)
}
