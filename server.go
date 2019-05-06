package goserve

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// serverConfig holds the information creating a server including
// the port to run the server on.
//
// GlobalMiddleware is a slice of Middleware that will be applied to
// every route.
type serverConfig struct {
	Port             int
	GracefulWait     time.Duration
	GlobalMiddleware []Middleware
}

// NewServerConfig returns a serverConfig struct with the
// default values.
func NewServerConfig() serverConfig {
	return serverConfig{
		Port:             8080,
		GracefulWait:     time.Second * 15,
		GlobalMiddleware: []Middleware{PanicRecover, LogHTTP},
	}
}

// NewServer creates and starts a new http.Server based on the provided
// configuration with the provided routes.
func NewServer(conf serverConfig, routes []Route) {
	r := newRouter(routes, conf.GlobalMiddleware)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", conf.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		log.Printf("Starting API on port %v...", conf.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), conf.GracefulWait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting down API...")
	os.Exit(0)
}
