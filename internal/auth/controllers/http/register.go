package controllers

import (
	"github.com/kapralovs/thinker/internal/auth"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoints(auth *echo.Group, uc auth.UseCase) {
	h := NewAuthHandler(uc)
	auth.POST("sign_in", h.SignIn)
	auth.POST("sign_up", h.SignUp)
}
