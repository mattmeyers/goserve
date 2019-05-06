package goserve

import "net/http"

// Route holds information for a single Gorilla/mux route.
// TODO: Add query parameters.
type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	Headers     []string
	HandlerFunc http.HandlerFunc
	Middleware  []Middleware
}
