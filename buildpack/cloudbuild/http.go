package cloudbuild

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := http.Server{
		Addr:        ":" + port,
		Handler:     nil,
		ReadTimeout: 10 * 60 * time.Second,
	}

	return server.ListenAndServe()
}
