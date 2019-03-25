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

// ServerConfig holds the information creating a server including
// the port to run the server on.
type ServerConfig struct {
	Port         int
	GracefulWait time.Duration
}

// NewServer creates and starts a new http.Server based on the provided
// configuration with the provided routes.
func NewServer(conf ServerConfig, routes []Route) {
	r := newRouter(routes)

	var p int
	if conf.Port <= 0 {
		p = 8080
	} else {
		p = conf.Port
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", p),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		log.Printf("Starting API on port %v...", p)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	var g time.Duration
	if conf.GracefulWait == 0 {
		g = time.Duration(15 * time.Second)
	} else {
		g = conf.GracefulWait
	}
	ctx, cancel := context.WithTimeout(context.Background(), g)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting down API...")
	os.Exit(0)
}
