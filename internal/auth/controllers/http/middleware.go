package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/kapralovs/thinker/internal/auth"
	"github.com/labstack/echo/v4"
)

type authMiddlewareHandler struct {
	usecase auth.UseCase
}

func NewAuthMiddlewareHandler(uc auth.UseCase) func(next echo.HandlerFunc) echo.HandlerFunc {
	return (&authMiddlewareHandler{usecase: uc}).AuthMiddleware
}

func (a *authMiddlewareHandler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeaderVal := c.Request().Header.Get("Authorization")
		if authHeaderVal == "" {
			return c.JSON(http.StatusUnauthorized, "auth failed")
		}

		token := authHeaderVal[len("Bearer "):]

		claims, err := a.usecase.ParseToken(token)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusUnauthorized, "auth failed!")
		}

		//create new context with token
		ctx := context.WithValue(c.Request().Context(), "token", claims)

		//clone request with a new context
		c.SetRequest(c.Request().Clone(ctx))

		//call next handler
		return next(c)
	}
}
