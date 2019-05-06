package goserve

import "net/http"

// Route holds information for a single Gorilla/mux route.
// TODO: Add query parameters.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
