# GoServe

GoServe wraps common HTTP server functionality into a simple library.  This library includes logging and panic recovery middleware for every provided `http.HandlerFunc`.  Additionally, a configurable graceful shutdown is provided.

# Installation

This library can be installed using `go get`:

```sh
go get -u github.com/mattmeyers/goserve
```

# Usage

In order to create a new server, use `goserve.NewServer`.  This method requires server configuration struct and a slice of route structs.  Currently, the server configuration struct allows for port and graceful shutdown configuration.

```go
type ServerConfig struct {
    Port         int            // Default: 8080
    GracefulWait time.Duration  // Default: 15 * time.Second
}
```

A route struct describes a `Gorilla/mux` route.

```go
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
```