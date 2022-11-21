package controllers

import "github.com/golang-jwt/jwt"

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	*AuthInfo
}

type TokenInfo struct {
	Token string `json:"token,omitempty"`
}

func generateToken() (string, error) {
	return token, nil
}

func parseToken() error {
	return nil
}
