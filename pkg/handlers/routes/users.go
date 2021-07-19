package routes

import "dashboard-api/pkg/controllers"

var userRoutes = []Route{
	{
		URI:    "/login",
		Method: "POST",
		Func:   controllers.Login,
		Auth:   false,
	},
	{
		URI:    "/register",
		Method: "POST",
		Func:   controllers.Register,
		Auth:   false,
	},
}
