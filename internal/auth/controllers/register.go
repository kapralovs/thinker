package controllers

import (
	"github.com/kapralovs/thinker/internal/auth"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoints(r echo.Echo, uc auth.Usecase) {
	h := NewAuthHandler(uc)
	r.POST("/auth/sign_in", h.SignIn)
	r.POST("/auth/sign_up", h.SignUp)
}