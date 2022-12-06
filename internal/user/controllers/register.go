package controllers

import (
	"github.com/kapralovs/thinker/internal/user"
	"github.com/labstack/echo/v4"
)

func RegisterEndpoints(r *echo.Echo, uc user.UseCase) {
	h := NewUserHandler(uc)
	user := r.Group("/user")
	user.GET("/get", h.GetUser)
	user.GET("/get_list", h.GetUsersList)
	user.POST("/edit", h.EditUser)
}
