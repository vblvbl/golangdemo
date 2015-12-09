package routers

import (
	"golangdemo/controllers"
	"golangdemo/core"
	"golangdemo/middlewares"
)

// getRouter returns the routers
func GetRouter() (router *core.Router) {
	router = core.NewRouter()

	// All routes go here
	router.HandleFunc("/", controllers.SayhelloName)

	//Static Controller
	router.PathPrefix("/").Handler(&controllers.Static{"/static/public", router})

	// All middlewares go here
	router.AddMiddleware("/", &middlewares.HTTPLogger{})

	return
}
