package controllers

import (
	"fmt"
	"goravel/app/models"
	"goravel/app/usecases"

	"github.com/golang-jwt/jwt/v5"
	"github.com/goravel/framework/contracts/http"
)

type AuthController struct {
	authUc usecases.AuthUsecase
}

type User struct {
	Username string
	Email    string
}

type CustomClaim struct {
	User
	jwt.RegisteredClaims
}

func NewAuthController(authUc usecases.AuthUsecase) *AuthController {
	return &AuthController{
		authUc: authUc,
	}
}

func (r *AuthController) Login(ctx http.Context) http.Response {
	request := new(models.LoginAtempt)
	if err := ctx.Request().Bind(request); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": err.Error(),
		})
	}
	fmt.Println(request)
	user, token, err := r.authUc.Login(ctx.Context(), request.Email, request.Password)
	if err != nil {
		if err.Error() == "invalid username or password" {
			return ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"error": err.Error(),
			})
		}
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"user":  user,
		"token": token,
	})
}

func (r *AuthController) Register(ctx http.Context) http.Response {
	request := new(models.User)
	if err := ctx.Request().Bind(request); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": err.Error(),
		})
	}
	user, token, err := r.authUc.Register(ctx.Context(), request)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"user":  user,
		"token": token,
	})
}
