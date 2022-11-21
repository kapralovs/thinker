package controllers

import (
	"github.com/kapralovs/thinker/internal/auth"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoints(r echo.Echo, uc auth.Usecase) {
	h := NewAuthHandler(uc)
	r.Get.
}
