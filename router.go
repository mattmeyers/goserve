package goserve

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter(routes []Route, middleware []Middleware) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = Chain(route.HandlerFunc, middleware...)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
