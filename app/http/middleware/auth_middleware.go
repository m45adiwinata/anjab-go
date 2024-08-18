package middleware

import (
	"fmt"

	"github.com/goravel/framework/contracts/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) CheckToken() http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Headers().Get("Authorization")
		if token == "" {
			ctx.Request().AbortWithStatusJson(http.StatusBadRequest, http.Json{
				"message": "you are not allowed",
			})
			return
		}
		fmt.Println(token)
		ctx.Request().Next()
	}
}
