package controllers

import (
	"github.com/kapralovs/thinker/internal/user"
	"github.com/labstack/echo/v4"
)

func RegisterHTTPEndpoints(router *echo.Echo, uc user.UseCase) {
	h := NewUserHandler(uc)

	users := router.Group("/users/")
	users.POST("create", h.CreateUser)
	users.PUT("edit/:id", h.EditUser)
	users.DELETE("delete/:id", h.DeleteUser)
	users.GET("get/:id", h.GetUser)
	users.GET("list", h.GetUsersList)
}
