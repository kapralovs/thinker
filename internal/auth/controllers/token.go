package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	*AuthInfo
}

type TokenInfo struct {
	Token string `json:"token,omitempty"`
}

func generateToken(a *AuthInfo, sString string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaimsExample{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"access",
		a,
	})

	return t.SignedString([]byte(sString))
}

func parseToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGN_STRING")), nil
	})
	if claims, ok := token.Claims.(*CustomClaimsExample); ok && token.Valid {
		fmt.Printf("%v %v\n", claims.Login, claims.ExpiresAt)
		fmt.Printf("%v %v\n", claims.Password, claims.ExpiresAt)
	} else {
		return err
	}

	return nil
}
