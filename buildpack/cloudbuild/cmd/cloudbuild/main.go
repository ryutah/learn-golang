package main

import (
	"log"
	"net/http"

	"github.com/ryutah/learn-golang/buildpack/cloudbuild"
)

func main() {
	log.Println("start server")

	if err := cloudbuild.Run(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
