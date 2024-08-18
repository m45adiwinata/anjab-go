package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
	"goravel/app/http/middleware"
)

func Api() {
	userController := controllers.NewUserController()
	authController := controllers.NewAuthController()
	authMiddleware := middleware.NewAuthMiddleware()
	facades.Route().Middleware(authMiddleware.CheckToken()).Get("/users/{id}", userController.Show)
	facades.Route().Post("/auth/login", authController.Login)
}
