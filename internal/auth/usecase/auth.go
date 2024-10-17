package usecase

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kapralovs/thinker/internal/auth"
	"github.com/kapralovs/thinker/internal/models"
)

type (
	authUseCase struct {
		repo auth.Repository
	}
)

func NewAuthUseCase(r auth.Repository) *authUseCase {
	return &authUseCase{repo: r}
}

func generateToken(a *models.AuthClaims, sString string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, a)

	return t.SignedString([]byte(sString))
}

func (uc *authUseCase) SignIn(username, password string) (string, error) {
	u, err := uc.repo.GetUser(username, password)
	if err != nil {
		return "", err
	}

	claims := &models.AuthClaims{
		User: u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Minute),
		},
	}

	token, err := generateToken(claims, os.Getenv("SIGN_STRING"))
	if err != nil {
		return "", err
	}

	u.CurrentToken = token

	if err = uc.repo.UpdateUser(u); err != nil {
		return "", err
	}

	return token, nil
}

func (uc *authUseCase) SignUp(username, password string) (string, error) {
	u := &models.User{Username: username, Password: password}

	claims := &models.AuthClaims{
		User: u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Hour),
		},
	}

	token, err := generateToken(claims, os.Getenv("SIGN_STRING"))
	if err != nil {
		return "", err
	}

	u.CurrentToken = token

	if err = uc.repo.CreateUser(u); err != nil {
		return "", fmt.Errorf("%s: %s", "can't create user", err.Error())
	}

	return token, nil
}

func (uc *authUseCase) ParseToken(tokenString string) (*models.AuthClaims, error) {
	tokenInfo, err := jwt.ParseWithClaims(tokenString, &models.AuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SIGN_STRING")), nil
		})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenInfo.Claims.(*models.AuthClaims)
	if ok && tokenInfo.Valid {
		u, err := uc.repo.GetUser(claims.User.Username, claims.User.Password)
		if err != nil {
			return nil, err
		}

		if u.CurrentToken != tokenString {
			return nil, errors.New("the token is expired")
		}
	}

	return claims, nil
}
