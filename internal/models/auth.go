package models

import "github.com/golang-jwt/jwt"

type AuthClaims struct {
	jwt.StandardClaims
	User *User `json:"user"`
}
