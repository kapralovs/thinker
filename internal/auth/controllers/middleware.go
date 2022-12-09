package controllers

import (
	"net/http"

	"github.com/kapralovs/thinker/internal/auth"
	"github.com/labstack/echo/v4"
)

type AuthMiddlewareHandler struct {
	usecase auth.UseCase
}

func NewAuthMiddlewareHandler(uc auth.UseCase) func(next echo.HandlerFunc) echo.HandlerFunc {
	return (&AuthMiddlewareHandler{usecase: uc}).AuthMiddleware
}

func (a *AuthMiddlewareHandler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeaderVal := c.Request().Header.Get("Authorization")
		if authHeaderVal == "" {
			return c.JSON(http.StatusUnauthorized, "auth failed")
		}

		token := authHeaderVal[len("Bearer "):]
		// fmt.Printf("Token: %s\n", token)

		err := a.usecase.ParseToken(token)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "auth failed!")
		}

		return next(c)
	}
}
