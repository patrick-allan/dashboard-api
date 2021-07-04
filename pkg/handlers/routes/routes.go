package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents api route
type Route struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
	Auth   bool
}

//defireRoutes adiciona as rotas no handler
func DefireRoutes(r *mux.Router) {
	routes := userRoutes
	for _, route := range routes {
		if route.Auth {
			/* r.HandleFunc(route.URI, middlewares.Auth(route.Func)).Methods(route.Method) */
			r.HandleFunc(route.URI, route.Func).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Func).Methods(route.Method)
		}
	}
}
