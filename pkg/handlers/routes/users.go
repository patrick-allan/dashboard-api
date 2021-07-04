package routes

import "dashboard-api/pkg/controllers"

var userRoutes = []Route{
	{
		URI:    "/login",
		Method: "POST",
		Func:   controllers.Login,
		Auth:   false,
	},
}
