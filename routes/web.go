package routes

import (
	"goravel/app/http/controllers"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	dashboardController := controllers.NewDashboardController()
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})
	facades.Route().Get("/dashboard", dashboardController.Index)
	facades.Route().Static("assets", "./resources/views/assets")
	facades.Route().Static("adminlte", "./public/adminlte")
	facades.Route().Static("local", "./public/local")
}
