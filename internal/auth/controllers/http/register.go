package controllers

import (
	"github.com/kapralovs/thinker/internal/auth"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoints(auth *echo.Group, uc auth.UseCase) {
	handler := NewAuthHandler(uc)

	auth.POST("sign_in", handler.signIn)
	auth.POST("sign_up", handler.signUp)
}
