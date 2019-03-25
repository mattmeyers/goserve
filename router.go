package goserve

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter(routes []Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logHTTP(panicRecover(handler))

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
