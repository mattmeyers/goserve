package goserve

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Middleware represents a function that wraps a http.HandlerFunc with
// additional functionality.
type Middleware mux.MiddlewareFunc

// Chain wraps a http.HandlerFunc with middleware functions.
func Chain(fn http.Handler, middleware ...Middleware) http.Handler {
	for _, m := range middleware {
		fn = m(fn)
	}
	return fn
}
