package controllers

import (
	"fmt"
	"net/http"

	uc "github.com/kapralovs/thinker/internal/auth/usecase"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeaderVal := c.Request().Header.Get("Authorization")
		if authHeaderVal == "" {
			fmt.Println("WITHOUT AUTH")
			return c.JSON(http.StatusUnauthorized, "auth failed")
		}

		token := authHeaderVal[:len("Bearer ")]
		fmt.Printf("Token: %s\n", token)

		err := uc.ParseToken(token)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "auth failed!")
		}

		return next(c)
	}
}
