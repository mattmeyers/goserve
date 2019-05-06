package goserve

import (
	"net/http"
)

// Middleware represents a function that wraps a http.HandlerFunc with
// additional functionality.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain wraps a http.HandlerFunc with middleware functions.
func Chain(fn http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	for _, m := range middleware {
		fn = m(fn)
	}
	return fn
}
