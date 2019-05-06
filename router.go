package goserve

import (
	"github.com/gorilla/mux"
)

func newRouter(routes []Route, globalMiddleware []mux.MiddlewareFunc) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(globalMiddleware...)
	for _, route := range routes {
		handler := Chain(route.HandlerFunc, route.Middleware...)

		router.
			Methods(route.Methods...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler).
			Headers(route.Headers...)
	}

	return router
}
