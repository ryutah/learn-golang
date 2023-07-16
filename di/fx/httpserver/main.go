package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			newHandler,
			newHTTPServer,
		),
		fx.Invoke(
			func(srv *http.Server) {},
		),
	).Run()
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	return mux
}

func newHTTPServer(lc fx.Lifecycle, handler http.Handler) (*http.Server, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return nil, err
	}
	srv := http.Server{
		Handler:           handler,
		ReadHeaderTimeout: 60 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := srv.Serve(lis); err != nil && err != http.ErrServerClosed {
					log.Printf("failed to serve: %v", err)
				}
			}()
			return nil
		},
		OnStop: srv.Shutdown,
	})

	return &srv, nil
}
