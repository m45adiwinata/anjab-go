package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
	"goravel/app/http/middleware"
	"goravel/app/repositories"
	"goravel/app/usecases"
)

func Api() {
	mysqlDDB := facades.Orm().Connection("mysql")
	authRepository := repositories.NewAuthRepository(mysqlDDB)
	authUsecase := usecases.NewAuthUsecase(authRepository)
	authController := controllers.NewAuthController(authUsecase)

	userController := controllers.NewUserController()
	authMiddleware := middleware.NewAuthMiddleware()
	facades.Route().Middleware(authMiddleware.CheckToken()).Get("/users/{id}", userController.Show)
	facades.Route().Prefix("/auth").Group(func(router route.Router) {
		router.Post("/login", authController.Login)
		router.Post("/register", authController.Register)
	})
}
