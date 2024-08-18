package controllers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AuthController struct {
}

type User struct {
	Username string
	Email    string
}

type CustomClaim struct {
	User
	jwt.RegisteredClaims
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (r *AuthController) Login(ctx http.Context) http.Response {
	config := facades.Config()
	request := ctx.Request().All()
	fmt.Println(request)
	jwtSecret := []byte(config.Env("JWT_SECRET").(string))
	claims := CustomClaim{
		User{
			Username: request["username"].(string),
			Email:    "dummy@gmail.com",
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		if err != nil {
			return ctx.Response().Json(http.StatusInternalServerError, http.Json{"message": "failed to get token string"})
		}
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"user":  request,
		"token": tokenString,
	})
}
